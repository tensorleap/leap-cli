#!/bin/bash
set -euo pipefail

# Get the machine name from the Azure metadata service
MECHINE_NAME=$(curl -s -H "Metadata:true" \
  "http://169.254.169.254/metadata/instance/compute/name?api-version=2021-02-01&format=text" \
  | sed 's/-.*//')

# Size of the image (GiB)
: "${SIZE_GB:=500}"

# Where your Azure Files (CIFS) share is mounted
SHARE_MNT="/mnt/batch/tasks/shared/LS_root/mounts/clusters/$MECHINE_NAME/code"

# Path to the image file that will live ON the share
IMG="$SHARE_MNT/tensorleap-volumes/$MECHINE_NAME.ext4"

# Local mountpoint for the loop-mounted ext4
MNT="/mnt/persist"
TLDATA_DIR="$MNT/tldata"

# --- 1) Wait for the CIFS share to be mounted ---
echo "Waiting for CIFS share at $SHARE_MNT ..."
for i in {1..120}; do
  if mountpoint -q "$SHARE_MNT"; then
    echo "Share mounted."
    break
  fi
  sleep 1
  if [[ $i -eq 120 ]]; then
    echo "ERROR: CIFS share $SHARE_MNT not mounted after 120s." >&2
    exit 1
  fi
done

# --- 2) Ensure the directory for the image exists (create on first run) ---
PARENT_DIR="$(dirname "$IMG")"
sudo mkdir -p "$PARENT_DIR"


# --- 3) Create the sparse image file + ext4 filesystem if missing (first run only) ---
if [[ ! -f "$IMG" ]]; then
  echo "Creating sparse ${SIZE_GB}GiB image at $IMG ..."
  sudo dd if=/dev/zero of="$IMG" bs=1M count=0 seek=$((SIZE_GB*1024))
  echo "Formatting $IMG as ext4 ..."
  sudo mkfs.ext4 -F "$IMG"
fi

# --- 4) Create and enable the systemd service ---
SERVICE_FILE="/etc/systemd/system/mnt-persist.service"
if [ ! -f "$SERVICE_FILE" ]; then
  echo "Creating systemd service for $MNT ..."

  sudo tee "$SERVICE_FILE" > /dev/null <<EOF
[Unit]
Description=Mount persistent ext4 image once Azure Batch share is ready
After=network-online.target remote-fs.target
Wants=network-online.target

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/usr/bin/bash -c 'for i in {1..120}; do \
  IMG="$IMG"; \
  MNT="$MNT"; \
  if [ -f "\$IMG" ]; then \
    echo "[$(date)] Found image, mounting..."; \
    mkdir -p "\$MNT"; \
    mount -o loop "\$IMG" "\$MNT" && \
    chmod 775 "\$MNT" && \
    chown azureuser:azureuser "\$MNT"; \
    echo "[$(date)] Mount and permissions done."; \
    exit 0; \
  fi; \
  echo "[$(date)] Waiting for Azure Batch share... (\$i)"; \
  sleep 2; \
done; \
echo "[$(date)] ERROR: image not found after 4 minutes."; exit 1'
ExecStop=/usr/bin/umount "$MNT"

[Install]
WantedBy=multi-user.target
EOF

  echo "Enabling and starting systemd service..."
  sudo systemctl daemon-reload
  sudo systemctl enable mnt-persist.service
  sudo systemctl start mnt-persist.service
else
  echo "Systemd service already exists — skipping creation."
fi

# wait for the mount to be ready
for i in {1..10}; do
  if mountpoint -q "$MNT"; then
    echo "Mount ready."
    break
  fi
  sleep 1
done

# --- 5) Set Docker data-root to $MNT/docker-data ---
# Ensure Docker starts only after the persist mount is ready (avoids disk-pressure and wrong storage on reboot)
ensure_docker_after_mount() {
  DOCKER_DROPIN_DIR="/etc/systemd/system/docker.service.d"
  DOCKER_DROPIN="$DOCKER_DROPIN_DIR/after-mnt-persist.conf"
  if [ ! -f "$DOCKER_DROPIN" ]; then
    echo "Configuring Docker to start after $MNT is mounted..."
    sudo mkdir -p "$DOCKER_DROPIN_DIR"
    sudo tee "$DOCKER_DROPIN" > /dev/null <<EOF
[Unit]
After=mnt-persist.service
Requires=mnt-persist.service
EOF
    sudo systemctl daemon-reload
  fi
}

update_docker_config() {
  DOCKER_DATA_ROOT="$MNT/docker-data"
  sudo mkdir -p "$DOCKER_DATA_ROOT"

  DOCKER_DAEMON_CONFIG=$(cat <<EOF
{
    "data-root": "$DOCKER_DATA_ROOT"
}
EOF
)

  DOCKER_DAEMON_CONFIG_PATH="/etc/docker/daemon.json"
  echo "$DOCKER_DAEMON_CONFIG" | sudo tee "$DOCKER_DAEMON_CONFIG_PATH" > /dev/null

  ensure_docker_after_mount

  sudo systemctl daemon-reload
  sudo systemctl restart docker
  sudo systemctl status docker --no-pager
}

# --- 7) Set TL_DISABLE_KUBE_STORAGE_TAINT globally (persists across reboots) ---
if ! grep -q '^TL_DISABLE_KUBE_STORAGE_TAINT=' /etc/environment 2>/dev/null; then
  echo "Setting TL_DISABLE_KUBE_STORAGE_TAINT=true in /etc/environment ..."
  echo 'TL_DISABLE_KUBE_STORAGE_TAINT=true' | sudo tee -a /etc/environment > /dev/null
fi
export TL_DISABLE_KUBE_STORAGE_TAINT=true

# --- 8) Verify and set permissions ---
if mountpoint -q "$MNT"; then
  echo "✅ Ready: $MNT is mounted (backed by $IMG on the CIFS share)."

  update_docker_config

  echo "Installing leap cli..."
  curl -s https://raw.githubusercontent.com/tensorleap/leap-cli/master/install.sh | bash
  echo "Leap cli installed."

  echo "The data directory is now available at $TLDATA_DIR. On first installation add --data-dir $TLDATA_DIR to your leap server install command."
  echo ""
  echo "If you use k3d and see node.kubernetes.io/disk-pressure taint (e.g. 'invalid capacity 0 on image filesystem'), remove it with:"
  echo "  kubectl taint nodes --all node.kubernetes.io/disk-pressure:NoSchedule-"
  echo "When creating a new k3d cluster, use only nodefs in eviction-hard so broken imagefs (0 capacity in k3d) does not trigger disk pressure:"
  echo "  k3d cluster create <name> --k3s-arg '--kubelet-arg=eviction-hard=nodefs.available<50G'"
else
  echo "⚠️  Warning: $MNT not mounted yet. It will mount automatically once the CIFS share is available."
fi
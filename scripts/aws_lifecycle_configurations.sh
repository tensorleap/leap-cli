#!/bin/bash

set -e

# Example of a script that can be used in AWS Lifecycle Configurations to install Tensorleap CLI

# Tensorleap directory

VOLUME_DIR="/home/ec2-user/SageMaker"
BASE_PATH="$VOLUME_DIR/.tensorleap"
mkdir -p "$BASE_PATH"
export TL_CLI_CONFIG_FILE="$BASE_PATH/cli-config.yaml"

# kubectl location
export KUBECONFIG="$BASE_PATH/kubeconfig"

TL_BIN_DIR="$BASE_PATH/bin"
mkdir -p "$TL_BIN_DIR"
export PATH="$PATH:$TL_BIN_DIR"

# Function to update Docker configuration and restart Docker service
update_docker_config() {
    # JSON configuration to be saved
    DOCKER_DAEMON_CONFIG=$(cat <<EOF
{
    "runtimes": {
        "nvidia": {
            "args": [],
            "path": "nvidia-container-runtime"
        }
    },
    "data-root": "$VOLUME_DIR/docker-data-root"
}
EOF
)

    # Path to Docker daemon configuration file
    DOCKER_DAEMON_CONFIG_PATH="/etc/docker/daemon.json"

    # Write the JSON configuration to the Docker daemon configuration file
    echo "$DOCKER_DAEMON_CONFIG" | sudo tee $DOCKER_DAEMON_CONFIG_PATH > /dev/null

    # Restart Docker to apply the changes
    sudo systemctl daemon-reload
    sudo systemctl restart docker

    # Check and display the status of Docker service
    sudo systemctl status docker --no-pager
}

ensure_leap_cli() {
    if ! command -v leap &> /dev/null; then
        echo "Tensorleap CLI not found. Installing..."
        curl -s https://raw.githubusercontent.com/tensorleap/leap-cli/master/install.sh | BIN_DIR=$TL_BIN_DIR bash
        
        if [ -e "$TL_CLI_CONFIG_FILE" ]; then
            # Change permissions of the config file
            chmod 644 "$TL_CLI_CONFIG_FILE"
        fi
    else
        echo "Tensorleap CLI is already installed."
    fi
}

update_docker_config
ensure_leap_cli

START_SHELL_SCRIPT=$(cat <<EOL
    export TL_CLI_CONFIG_FILE="$TL_CLI_CONFIG_FILE"
    export TL_BIN_DIR="$TL_BIN_DIR"
    export KUBECONFIG="$KUBECONFIG"
    alias k="leap server tools kubectl -n tensorleap"
EOL
)

# escape the $PATH variable
START_SHELL_SCRIPT_2=$(cat <<'EOL'
    export PATH="$PATH:$TL_BIN_DIR"
EOL
)


# Only .profile is run on login shell (not .bashrc or .bash_profile)
PROFILE_FILE_PATH="/home/ec2-user/.profile"
if [ ! -f "$PROFILE_FILE_PATH" ]; then
    echo "Creating .profile file..."
    touch "$PROFILE_FILE_PATH"
    chmod 644 "$PROFILE_FILE_PATH"
fi
echo "$START_SHELL_SCRIPT" >> "$PROFILE_FILE_PATH"
echo "$START_SHELL_SCRIPT_2" >> "$PROFILE_FILE_PATH"


# Create a script to upgrade the Tensorleap CLI
UPGRADE_LEAP_SCRIPT=$(cat <<EOL
#!/bin/bash
leap cli upgrade -s | BIN_DIR=$TL_BIN_DIR bash
EOL
)
UPGRADE_LEAP_SCRIPT_PATH="$TL_BIN_DIR/upgrade_leap"
echo "$UPGRADE_LEAP_SCRIPT" > "$UPGRADE_LEAP_SCRIPT_PATH"
chmod +x "$UPGRADE_LEAP_SCRIPT_PATH"

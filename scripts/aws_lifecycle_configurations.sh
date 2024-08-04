#!/bin/bash

set -e

# Example of a script that can be used in AWS Lifecycle Configurations to install Tensorleap CLI

# Tensorleap directory
BASE_PATH="/home/ec2-user/SageMaker/tensorleap"
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
    DOCKER_DAEMON_CONFIG='{
        "runtimes": {
            "nvidia": {
                "args": [],
                "path": "nvidia-container-runtime"
            }
        },
        "data-root":"/home/ec2-user/SageMaker/docker-data-root"
    }'

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
            chmod 600 "$TL_CLI_CONFIG_FILE"
        fi
    else
        echo "Tensorleap CLI is already installed."
    fi
}

update_docker_config
ensure_leap_cli

BASHRC_CONTENT=$(cat <<EOL
    export TL_CLI_CONFIG_FILE="$TL_CLI_CONFIG_FILE"
    export TL_BIN_DIR="$TL_BIN_DIR"
    export KUBECONFIG="$KUBECONFIG"
    alias k="leap server tools kubectl -n tensorleap"
EOL
)

BASHRC_CONTENT_2=$(cat <<'EOL'
    export PATH="$PATH:$TL_BIN_DIR"
EOL
)


BASHRC_FILE="/home/ec2-user/.bashrc"
echo "$BASHRC_CONTENT" >> "$BASHRC_FILE"
echo "$BASHRC_CONTENT_2" >> "$BASHRC_FILE"


# Create a script to upgrade the Tensorleap CLI
UPGRADE_LEAP_SCRIPT=$(cat <<EOL
#!/bin/bash
leap cli upgrade -s | BIN_DIR=$TL_BIN_DIR bash
EOL
)
UPGRADE_LEAP_SCRIPT_PATH="$TL_BIN_DIR/upgrade_leap"
echo "$UPGRADE_LEAP_SCRIPT" > "$UPGRADE_LEAP_SCRIPT_PATH"

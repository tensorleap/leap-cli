#!/bin/bash

set -e
curl -s https://raw.githubusercontent.com/tensorleap/leap-cli/master/scripts/aws_lifecycle_configurations.sh | sudo bash

# Ignore ShellCheck warning SC1090 for non-constant source path
# shellcheck source=/dev/null
source ~/.profile
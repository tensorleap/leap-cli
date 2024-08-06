#!/bin/bash

set -e
curl -s https://raw.githubusercontent.com/tensorleap/leap-cli/master/scripts/aws_lifecycle_configurations.sh | sudo bash
source ~/.profile
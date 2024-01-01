#!/usr/bin/env bash

# Get two arguments: image name and branch name - by default it will use master branch

set -e

DOCKER_REGISTRY=us-central1-docker.pkg.dev/tensorleap/main
NODE_SERVER_REPO=node-server

NODE_SERVER_BRANCH=${2:-master}
function get_image_name() {
    IMAGE_NAME=${DOCKER_REGISTRY}/${NODE_SERVER_REPO}
    LAST_COMMIT_HASH=$(gh api repos/tensorleap/$NODE_SERVER_REPO/commits/$NODE_SERVER_BRANCH --jq '.sha')
    IMAGE_TAG=$NODE_SERVER_BRANCH-$(echo $LAST_COMMIT_HASH | cut -c 1-8)
    IMAGE_BUILDER_TAG=${IMAGE_TAG}-builder
    IMAGE=$IMAGE_NAME:$IMAGE_BUILDER_TAG
    echo $IMAGE
}

IMAGE=$1
if [ -z "$IMAGE" ]; then
    IMAGE=$(get_image_name)
fi

echo "Pulling node server image: ${IMAGE}";
docker pull ${IMAGE}

echo "Removing old server api...";
rm -rf ./pkg/tensorleapapi

echo "Generating server api...";
docker run --rm -v $(pwd)/pkg/tensorleapapi:/usr/app/generated/tensorleapapi ${IMAGE} sh -c "npm run generate-go-client"

echo "Formatting code..."
make fmt


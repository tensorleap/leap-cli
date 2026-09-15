#!/usr/bin/env bash

# Get two arguments: image name and branch name - by default it will use master branch

set -e

DOCKER_REGISTRY=public.ecr.aws/tensorleap
NODE_SERVER_REPO=node-server

NODE_SERVER_BRANCH=${2:-master}
function get_image_name() {
    IMAGE_NAME=${DOCKER_REGISTRY}/${NODE_SERVER_REPO}
    LAST_COMMIT_HASH="$(gh api "repos/tensorleap/$NODE_SERVER_REPO/commits/$NODE_SERVER_BRANCH" --jq '.sha')"
    IMAGE_TAG="$NODE_SERVER_BRANCH-$(echo "$LAST_COMMIT_HASH" | cut -c 1-8)"
    IMAGE_BUILDER_TAG=${IMAGE_TAG}-builder
    IMAGE=$IMAGE_NAME:$IMAGE_BUILDER_TAG
    echo "$IMAGE"
}

IMAGE=$1
if [ -z "$IMAGE" ]; then
    IMAGE=$(get_image_name)
fi

if ! docker image inspect "${IMAGE}" >/dev/null 2>&1; then
    # The builder image is never published; build it from a node-server checkout.
    echo "Building node server builder image: ${IMAGE} (branch ${NODE_SERVER_BRANCH})";
    SRC_DIR=$(mktemp -d)
    trap 'rm -rf "${SRC_DIR}"' EXIT
    git clone --quiet --depth 1 --branch "${NODE_SERVER_BRANCH}" git@github.com:tensorleap/node-server.git "${SRC_DIR}"
    docker build --target builder --build-arg NPM_TOKEN="${NPM_TOKEN:?set NPM_TOKEN to build the node-server builder image}" -t "${IMAGE}" "${SRC_DIR}"
fi

echo "Removing old server api...";
rm -rf ./pkg/tensorleapapi

echo "Generating server api...";
docker run --rm -v "$(pwd)/pkg/tensorleapapi:/usr/app/generated/tensorleapapi" "${IMAGE}" sh -c "npm run generate-go-client"

echo "Formatting code..."
make fmt


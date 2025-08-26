#!/usr/bin/env bash

# This is based on k3d installation script: https://github.com/k3d-io/k3d/blob/main/install.sh

APP_NAME="leap"
REPO_URL="https://github.com/tensorleap/leap-cli"

: "${USE_SUDO:="true"}"
: "${BIN_DIR:="/usr/local/bin"}"

# initArch discovers the architecture for this system.
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
    armv5*) ARCH="armv5";;
    armv6*) ARCH="armv6";;
    armv7*) ARCH="arm";;
    aarch64) ARCH="arm64";;
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
  esac
}

# initOS discovers the operating system for this system.
initOS() {
  OS=$(uname|tr '[:upper:]' '[:lower:]')

  case "$OS" in
    # Minimalist GNU for Windows
    mingw*)
      OS="windows"
      USE_SUDO="false"
      if [[ ! -d "$BIN_DIR" ]]; then
        # mingw bash that ships with Git for Windows doesn't have /usr/local/bin but ~/bin is first entry in the path
        mkdir -p ~/bin
        BIN_DIR=~/bin
      fi
      ;;
  esac
}

# runs the given command as root (detects if we are root already)
runAsRoot() {
  local CMD="$*"

  if [ $EUID -ne 0 ] && [ "$USE_SUDO" = "true" ]; then
    CMD="sudo $CMD"
  fi

  $CMD
}

# scurl invokes `curl` with secure defaults
scurl() {
  # - `--proto =https` requires that all URLs use HTTPS. Attempts to call http://
  #   URLs will fail.
  # - `--tlsv1.2` ensures that at least TLS v1.2 is used, disabling less secure
  #   prior TLS versions.
  # - `--fail` ensures that the command fails if HTTP response is not 2xx.
  # - `--show-error` causes curl to output error messages when it fails (when
  #   also invoked with -s|--silent).
  curl --proto "=https" --tlsv1.2 --fail --show-error "$@"
}

# verifySupported checks that the os/arch combination is supported for
# binary builds.
verifySupported() {
  local supported="linux-386\nlinux-amd64\nlinux-arm\nlinux-arm64\ndarwin-amd64\ndarwin-arm64"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuilt binary for ${OS}-${ARCH}."
    echo "To build from source, go to $REPO_URL"
    exit 1
  fi

  if ! type "curl" > /dev/null && ! type "wget" > /dev/null; then
    echo "Either curl or wget is required"
    exit 1
  fi
}

# checkInstalledVersion checks which version is installed and if it needs to be changed.
checkInstalledVersion() {
  local version
  if [[ -f "${BIN_DIR}/${APP_NAME}" ]]; then
    version=$(leap --version)
    if [[ "$version" == "$TAG" ]]; then
      echo "leap ${version} is already ${DESIRED_VERSION:-latest}"
      return 0
    else
      echo "leap ${TAG} is available. Changing from version ${version}."
      return 1
    fi
  else
    return 1
  fi
}

# checkTagProvided checks whether TAG has provided as an environment variable so we can skip checkLatestVersion.
checkTagProvided() {
  [[ -n "$TAG" ]]
}

# checkLatestVersion grabs the latest version string from the releases
checkLatestVersion() {
  local latest_release_url="$REPO_URL/releases/latest"
  if type "curl" > /dev/null; then
    TAG=$(scurl -Ls -o /dev/null -w "%{url_effective}" $latest_release_url | grep -oE "[^/]+$" )
  elif type "wget" > /dev/null; then
    TAG=$(wget $latest_release_url --server-response -O /dev/null 2>&1 | awk '/^\s*Location: /{DEST=$2} END{ print DEST}' | grep -oE "[^/]+$")
  fi
}

# downloadFile downloads the latest binary package and also the checksum
# for that binary.
downloadFile() {
  DIST_FILE="leap-$OS-$ARCH"
  DOWNLOAD_URL="$REPO_URL/releases/download/$TAG/$DIST_FILE"
  # if [[ "$OS" == "windows" ]]; then
  #   DOWNLOAD_URL=${DOWNLOAD_URL}.exe
  # fi
  TMP_ROOT="$(mktemp -dt leap-binary-XXXXXX)"
  TMP_FILE="$TMP_ROOT/$DIST_FILE"
  if type "curl" > /dev/null; then
    scurl -sL "$DOWNLOAD_URL" -o "$TMP_FILE"
  elif type "wget" > /dev/null; then
    wget -q -O "$TMP_FILE" "$DOWNLOAD_URL"
  fi
}

# installFile verifies the SHA256 for the file, then unpacks and
# installs it.
installFile() {
  echo "Preparing to install $APP_NAME into ${BIN_DIR}"
  runAsRoot chmod +x "$TMP_FILE"
  runAsRoot cp "$TMP_FILE" "$BIN_DIR/$APP_NAME"
  echo "$APP_NAME installed into $BIN_DIR/$APP_NAME"
}

# fail_trap is executed if an error occurs.
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    # Track CLI installation failed
    trackEvent "cli_install_failed" "\"bin_dir\":\"$BIN_DIR\",\"use_sudo\":\"$USE_SUDO\",\"exit_code\":\"$result\""
    
    if [[ -n "$INPUT_ARGUMENTS" ]]; then
      echo "Failed to install $APP_NAME with the arguments provided: $INPUT_ARGUMENTS"
      help
    else
      echo "Failed to install $APP_NAME"
    fi
    echo -e "\tFor support, go to $REPO_URL."
  fi
  cleanup
  exit $result
}

# testVersion tests the installed client to make sure it is working.
testVersion() {
  if ! command -v $APP_NAME &> /dev/null; then
    # shellcheck disable=SC2016
    echo "$APP_NAME not found. Is $BIN_DIR on your "'$PATH?'
    exit 1
  fi
  echo "Run '$APP_NAME --help' to see what you can do with it."
}

# help provides possible cli installation arguments
help () {
  echo "Accepted cli arguments are:"
  echo -e "\t[--help|-h ] ->> prints this help"
  echo -e "\t[--no-sudo]  ->> install without sudo"
}

# trackEvent sends an analytics event to track CLI installation
trackEvent() {
  local event_type="$1"
  local properties="$2"
  
  # Only track if curl is available and we're not in a test environment
  if type "curl" > /dev/null && [[ -z "${TESTING:-}" ]]; then
    # Detect AWS environment
    local aws_env="false"
    if [[ -n "${AWS_EXECUTION_ENV:-}" ]] || [[ -f "/sys/hypervisor/uuid" ]] || [[ -f "/var/lib/cloud/instance" ]]; then
      aws_env="true"
    fi
    
    # Get current username
    local current_user=$(whoami 2>/dev/null || echo 'unknown')
    
    # Prepare the event data
    local event_data="{\"event\":\"$event_type\",\"properties\":{\"token\":\"f1bf46fb339d8c2930cde8c1acf65491\",\"time\":$(date +%s),\"os\":\"$(uname | tr '[:upper:]' '[:lower:]')\",\"arch\":\"$(uname -m)\",\"timestamp\":\"$(date -u +%Y-%m-%dT%H:%M:%SZ)\",\"distinct_id\":\"$current_user\",\"whoami\":\"$current_user\",\"aws_environment\":\"$aws_env\""
    
    # Add custom properties if provided
    if [[ -n "$properties" ]]; then
      event_data="$event_data,$properties"
    fi
    
    event_data="$event_data}}"
    
    # Send the event to Mixpanel (non-blocking)
    curl -s -X POST "https://api.mixpanel.com/track" \
      -d "data=$(echo "$event_data" | base64)" \
      --max-time 5 \
      --connect-timeout 3 \
      >/dev/null 2>&1 &
  fi
}

# cleanup temporary files
cleanup() {
  if [[ -d "${TMP_ROOT:-}" ]]; then
    rm -rf "$TMP_ROOT"
  fi
}

# Execution

#Stop execution on any error
trap "fail_trap" EXIT
set -e

# Parsing input arguments (if any)
# TODO: fix this shellcheck instead of disabling it
# shellcheck disable=SC2124
export INPUT_ARGUMENTS="${@}"
set -u
while [[ $# -gt 0 ]]; do
  case $1 in
    '--no-sudo')
       USE_SUDO="false"
       ;;
    '--help'|-h)
       help
       exit 0
       ;;
    *) exit 1
       ;;
  esac
  shift
done
set +u

initArch
initOS
verifySupported

# Track CLI installation started
trackEvent "cli_install_started" "\"bin_dir\":\"$BIN_DIR\",\"use_sudo\":\"$USE_SUDO\""

checkTagProvided || checkLatestVersion
if ! checkInstalledVersion; then
  downloadFile
  installFile
fi
testVersion

# Track CLI installation success
trackEvent "cli_install_success" "\"bin_dir\":\"$BIN_DIR\",\"use_sudo\":\"$USE_SUDO\""

cleanup

leap -h

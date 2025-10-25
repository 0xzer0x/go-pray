#!/usr/bin/env bash
set -euo pipefail

# colors
__CDEF="\033[0m"      # default color
__CCIN="\033[0;36m"   # info color
__CGSC="\033[0;32m"   # success color
__CRER="\033[0;31m"   # error color
__CWAR="\033[0;33m"   # waring color
__b_CDEF="\033[1;37m" # bold default color
__b_CCIN="\033[1;36m" # bold info color
__b_CGSC="\033[1;32m" # bold success color
__b_CRER="\033[1;31m" # bold error color
__b_CWAR="\033[1;33m" # bold warning color

# constants
__RELEASE_WORKFLOW_FILE="release.yml"
__GITHUB_OIDC_ISSUER="https://token.actions.githubusercontent.com"
__BASE_INSTALL_URL="https://github.com/0xzer0x/go-pray/releases/download/"
__OS="$(uname -o)"
__ARCH="$(uname -m)"
__DISTRO=""
_RELEASE_URL=""
_TEMPDIR=""

# script vars
INSTALL_VERSION="${INSTALL_VERSION:-latest}"
INSTALL_DIR="${INSTALL_DIR:-./bin}"

__prompt() {
  case ${1} in
  "-s" | "--success")
    echo -e "${__b_CGSC}[$] ${2}${__CDEF}"
    ;; # print success message
  "-e" | "--error")
    echo -e "${__b_CRER}[!] ${2}${__CDEF}"
    ;; # print error message
  "-w" | "--warning")
    echo -e "${__b_CWAR}[-] ${2}${__CDEF}"
    ;; # print warning message
  "-i" | "--info")
    echo -e "${__b_CCIN}[+] ${2}${__CDEF}"
    ;; # print info message
  *)
    echo -e "$@"
    ;;
  esac
}

__create-tempdir() {
  _TEMPDIR="$(mktemp -d)"
}

_cleanup() {
  if [ -d "${_TEMPDIR}" ]; then
    __prompt -w "Removing temporary download directory"
    rm -r "${_TEMPDIR}"
  fi
}

__validate_platform() {
  local _required_commands=(jq curl)
  if ! (command -v apt || command -v dnf || command -v apk) &>/dev/null; then
    _required_commands+=(cosign tar)
  fi

  # NOTE: check compatible os/arch
  if ! grep -i linux <<<"${__OS}" &>/dev/null || [ "${__ARCH}" != "x86_64" ]; then
    __prompt -e "Platform not supported, currently only linux/amd64 platform is supported"
    exit 1
  fi

  # NOTE: check required commands for script
  for cmd in "${_required_commands[@]}"; do
    if ! command -v "${cmd}" &>/dev/null; then
      __prompt -e "Required command not found: ${cmd}"
      exit 1
    fi
  done
}

__detect_distro() {
  if command -v apt &>/dev/null; then
    __DISTRO="debian"
  elif command -v dnf &>/dev/null; then
    __DISTRO="fedora"
  elif command -v apk &>/dev/null; then
    __DISTRO="alpine"
  elif (command -v yay || command -v paru) &>/dev/null; then
    __DISTRO="archlinux"
  elif command -v nixos-rebuild &>/dev/null; then
    __DISTRO="nixos"
  else
    __DISTRO="unknown"
  fi
}

_set_install_version() {
  if [ "${INSTALL_VERSION}" = "latest" ]; then
    __prompt -i "Fetching latest release version"
    INSTALL_VERSION="$(curl -sf https://api.github.com/repos/0xzer0x/go-pray/releases/latest | jq -r '.tag_name')"
  fi

  if [ "${__DISTRO}" = "archlinux" ]; then
    __prompt -w "Specifying custom INSTALL_VERSION is not supported on Archlinux"
  fi
  __prompt -i "Installing go-pray version: ${INSTALL_VERSION#v}"
  _RELEASE_URL="${__BASE_INSTALL_URL}/v${INSTALL_VERSION#v}"
}

_download_archive() {
  local _archive_name="go-pray_linux_x86_64.tar.gz"

  __create-tempdir
  __prompt -i "Downloading files in temp directory: ${_TEMPDIR}"

  __prompt -i "Downloading ${_archive_name}"
  curl -fLo "${_TEMPDIR}/${_archive_name}" "${_RELEASE_URL}/${_archive_name}"

  __prompt -i "Extracting tar archive"
  tar -C "${_TEMPDIR}" -xvzf "${_TEMPDIR}/${_archive_name}"
}

_verify_signature() {
  local _identity="https://github.com/0xzer0x/go-pray/.github/workflows/${__RELEASE_WORKFLOW_FILE}@refs/tags/v${INSTALL_VERSION#v}"
  local _signature_file="go-pray_linux_amd64.sig"

  if [ ! -r "${_TEMPDIR}/go-pray" ]; then
    __prompt -e "go-pray binary does not exist"
    exit 1
  fi

  __prompt -i "Fetching signature bundle"
  curl -fLo "${_TEMPDIR}/${_signature_file}" "${_RELEASE_URL}/${_signature_file}"

  __prompt -i "Verifying binary signature"
  cosign verify-blob \
    --bundle="${_TEMPDIR}/${_signature_file}" \
    --certificate-identity="${_identity}" \
    --certificate-oidc-issuer="${__GITHUB_OIDC_ISSUER}" \
    "${_TEMPDIR}/go-pray"
}

_install_binary() {
  if [ ! -r "${_TEMPDIR}/go-pray" ]; then
    __prompt -e "go-pray binary does not exist"
    exit 1
  elif [ ! -d "${INSTALL_DIR}" ]; then
    __prompt -w "Creating install directory: ${INSTALL_DIR}"
    mkdir -p "${INSTALL_DIR}"
  fi

  mv "${_TEMPDIR}/go-pray" "${INSTALL_DIR}/"
  __prompt -s "Successfully installed ${INSTALL_DIR}/go-pray"
}

_install_debian() {
  local _package_name="go-pray_${INSTALL_VERSION#v}_linux_amd64.deb"

  __create-tempdir
  __prompt -i "Downloading files in temp directory: ${_TEMPDIR}"

  __prompt -i "Downloading ${_package_name}"
  curl -fLo "${_TEMPDIR}/${_package_name}" "${_RELEASE_URL}/${_package_name}"

  __prompt -i "Installing ${_package_name}"
  if [ "$(id -u)" = "0" ]; then
    apt-get install "${_TEMPDIR}/${_package_name}"
  else
    sudo apt-get install "${_TEMPDIR}/${_package_name}"
  fi
}

_install_fedora() {
  local _package_name="go-pray_${INSTALL_VERSION#v}_linux_amd64.rpm"

  __create-tempdir
  __prompt -i "Downloading files in temp directory: ${_TEMPDIR}"

  __prompt -i "Downloading ${_package_name}"
  curl -fLo "${_TEMPDIR}/${_package_name}" "${_RELEASE_URL}/${_package_name}"

  __prompt -i "Installing ${_package_name}"
  if [ "$(id -u)" = "0" ]; then
    dnf install "${_TEMPDIR}/${_package_name}"
  else
    sudo dnf install "${_TEMPDIR}/${_package_name}"
  fi
}

_install_alpine() {
  local _package_name="go-pray_${INSTALL_VERSION#v}_linux_amd64.rpm"

  __create-tempdir
  __prompt -i "Downloading files in temp directory: ${_TEMPDIR}"

  __prompt -i "Downloading ${_package_name}"
  curl -fLo "${_TEMPDIR}/${_package_name}" "${_RELEASE_URL}/${_package_name}"

  __prompt -i "Installing ${_package_name}"
  if [ "$(id -u)" = "0" ]; then
    apk add "${_TEMPDIR}/${_package_name}"
  else
    sudo apk add "${_TEMPDIR}/${_package_name}"
  fi
}

_install_archlinux() {
  local _package_name="go-pray-bin"

  __prompt -w "Installing ${_package_name} from AUR"
  if command -v yay &>/dev/null; then
    yay -S "${_package_name}"
  else
    paru -S "${_package_name}"
  fi
}

_install_nixos() {
  __prompt -w "Automated install for go-pray is not supported on NixOS :("
  __prompt -w "Add the following nixpkg to your overlays: https://github.com/0xzer0x/infra/blob/main/pkgs/go-pray/default.nix"
}

_install_generic() {
  _download_archive
  _verify_signature
  _install_binary
}

_install() {
  case "${__DISTRO}" in
  debian)
    _install_debian
    ;;
  fedora)
    _install_fedora
    ;;
  alpine)
    _install_alpine
    ;;
  archlinux)
    _install_archlinux
    ;;
  nixos)
    _install_nixos
    ;;
  *)
    _install_generic
    ;;
  esac
}

_main() {
  __validate_platform
  __detect_distro
  _set_install_version
  _install
  _cleanup
}

_main

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

__cleanup() {
  __prompt -w "removing temporary download directory"
  rm -r "${_TEMPDIR}"
}

__validate-platform() {
  # NOTE: check compatible os/arch
  if ! grep -i linux <<<"${__OS}" &>/dev/null || [ "${__ARCH}" != "x86_64" ]; then
    __prompt -e "platform not supported, currently only linux/amd64 platform is supported"
    exit 1
  fi

  # NOTE: check required commands for script
  for cmd in cosign jq curl tar; do
    if ! command -v "${cmd}" &>/dev/null; then
      __prompt -e "required command not found: ${cmd}"
      exit 1
    fi
  done
}

_set-install-version() {
  if [ "${INSTALL_VERSION}" = "latest" ]; then
    __prompt -i "fetching latest release version"
    INSTALL_VERSION="$(curl -sf https://api.github.com/repos/0xzer0x/go-pray/releases/latest | jq -r '.tag_name')"
  fi
  __prompt -i "installing go-pray version: ${INSTALL_VERSION#v}"
  _RELEASE_URL="${__BASE_INSTALL_URL}/v${INSTALL_VERSION#v}"
}

_download-archive() {
  local _archive_name="go-pray_linux_x86_64.tar.gz"

  __create-tempdir
  __prompt -i "downloading files in temp directory: ${_TEMPDIR}"

  __prompt -i "downloading ${_archive_name}"
  curl -fLo "${_TEMPDIR}/${_archive_name}" "${_RELEASE_URL}/${_archive_name}"

  __prompt -i "extracting tar archive"
  tar -C "${_TEMPDIR}" -xvzf "${_TEMPDIR}/${_archive_name}"
}

_verify-signature() {
  local _identity="https://github.com/0xzer0x/go-pray/.github/workflows/${__RELEASE_WORKFLOW_FILE}@refs/tags/${INSTALL_VERSION}"
  local _signature_file="go-pray_linux_amd64.sig"

  if [ ! -r "${_TEMPDIR}/go-pray" ]; then
    __prompt -e "go-pray binary does not exist"
    exit 1
  fi

  __prompt -i "fetching signature bundle"
  curl -fLo "${_TEMPDIR}/${_signature_file}" "${_RELEASE_URL}/${_signature_file}"

  __prompt -i "verifying binary signature"
  cosign verify-blob \
    --bundle="${_TEMPDIR}/${_signature_file}" \
    --certificate-identity="${_identity}" \
    --certificate-oidc-issuer="${__GITHUB_OIDC_ISSUER}" \
    "${_TEMPDIR}/go-pray"
}

_install-binary() {
  if [ ! -r "${_TEMPDIR}/go-pray" ]; then
    __prompt -e "go-pray binary does not exist"
    exit 1
  elif [ ! -d "${INSTALL_DIR}" ]; then
    __prompt -w "creating install directory: ${INSTALL_DIR}"
    mkdir -p "${INSTALL_DIR}"
  fi

  mv "${_TEMPDIR}/go-pray" "${INSTALL_DIR}/"
}

_main() {
  __validate-platform
  _set-install-version
  _download-archive
  _verify-signature
  _install-binary
  __cleanup
  __prompt -s "successfully installed ${INSTALL_DIR}/go-pray"
}

_main

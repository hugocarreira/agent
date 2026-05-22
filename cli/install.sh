#!/usr/bin/env bash
# agentrc — One-line installer (macOS, Linux, WSL)
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/hugocarreira/agentrc/main/cli/install.sh | bash
#
# Override release URL:
#   INSTALL_URL="https://github.com/hugocarreira/agentrc/releases/latest/download" \
#     curl -fsSL ... | bash
set -euo pipefail

REPO="${REPO:-hugocarreira/agentrc}"
INSTALL_DIR="${INSTALL_DIR:-${HOME}/.local/bin}"
RELEASES_URL="${INSTALL_URL:-https://github.com/${REPO}/releases/latest/download}"

# ── Colors ───────────────────────────────────────────────────────────
RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; CYAN='\033[0;36m'; NC='\033[0m'
info()  { echo -e "${CYAN}ℹ${NC}  $*"; }
ok()    { echo -e "${GREEN}✓${NC}  $*"; }
warn()  { echo -e "${YELLOW}⚠${NC}  $*"; }
die()   { echo -e "${RED}✗${NC}  $*" >&2; exit 1; }

# ── Detect platform ──────────────────────────────────────────────────
detect_platform() {
	local os arch
	case "$(uname -s)" in
		Linux)  os="linux" ;;
		Darwin) os="darwin" ;;
		MINGW*|MSYS*|CYGWIN*) os="windows" ;;
		*) die "Unsupported OS: $(uname -s)" ;;
	esac
	case "$(uname -m)" in
		x86_64|amd64) arch="amd64" ;;
		arm64|aarch64) arch="arm64" ;;
		*) die "Unsupported architecture: $(uname -m)" ;;
	esac
	echo "${os}-${arch}"
}

main() {
	echo ""
	echo -e "${CYAN}  agentrc — installer${NC}"
	echo -e "  ─────────────────────────────────────────────────────────"
	echo ""

	local platform binary_name url tmpfile

	platform=$(detect_platform)
	binary_name="agentrc-${platform}"

	url="${RELEASES_URL}/${binary_name}"
	info "Platform : ${platform}"
	info "Download : ${url}"

	tmpfile=$(mktemp)
	trap "rm -f ${tmpfile}" EXIT

	if command -v curl &>/dev/null; then
		curl -fsSL --progress-bar -o "${tmpfile}" "${url}" || die "Download failed: ${url}"
	elif command -v wget &>/dev/null; then
		wget -q --show-progress -O "${tmpfile}" "${url}" || die "Download failed: ${url}"
	else
		die "Neither curl nor wget found."
	fi
	ok "Downloaded $(du -h "${tmpfile}" | cut -f1)"

	# macOS: strip quarantine
	if [[ "$(uname -s)" == "Darwin" ]]; then
		xattr -d com.apple.quarantine "${tmpfile}" 2>/dev/null || true
	fi

	chmod +x "${tmpfile}"
	mkdir -p "${INSTALL_DIR}"
cp "${tmpfile}" "${INSTALL_DIR}/agentrc"
ok "Installed → ${INSTALL_DIR}/agentrc"

	# PATH warning
	if ! echo "${PATH}" | tr ':' '\n' | grep -qxF "${INSTALL_DIR}"; then
		echo ""
		warn "${INSTALL_DIR} is not in your PATH. Add:"
		echo ""
		echo -e "    ${GREEN}export PATH=\"\${HOME}/.local/bin:\${PATH}\"${NC}"
		echo ""
	fi

	echo ""
	echo -e "${GREEN}  ✓ Installation complete!${NC}"
	echo ""
	echo "  Run: agentrc setup"
	echo ""
}

main "$@"

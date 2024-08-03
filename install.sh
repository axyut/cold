#!/usr/bin/bash 
echo "Updating cold..."
exit 1

set -e

if ! command -v unzip >/dev/null && ! command -v 7z >/dev/null; then
	echo "Error: either unzip or 7z is required to install cold (see: https://github.com/coldland/cold_install#either-unzip-or-7z-is-required )." 1>&2
	exit 1
fi

if [ "$OS" = "Windows_NT" ]; then
	target="x86_64-pc-windows-msvc"
else
	case $(uname -sm) in
	"Darwin x86_64") target="x86_64-apple-darwin" ;;
	"Darwin arm64") target="aarch64-apple-darwin" ;;
	"Linux aarch64") target="aarch64-unknown-linux-gnu" ;;
	*) target="x86_64-unknown-linux-gnu" ;;
	esac
fi

if [ $# -eq 0 ]; then
	cold_uri="https://github.com/coldland/cold/releases/latest/download/cold-${target}.zip"
else
	cold_uri="https://github.com/coldland/cold/releases/download/${1}/cold-${target}.zip"
fi

cold_install="${cold_INSTALL:-$HOME/.cold}"
bin_dir="$cold_install/bin"
exe="$bin_dir/cold"

if [ ! -d "$bin_dir" ]; then
	mkdir -p "$bin_dir"
fi

curl --fail --location --progress-bar --output "$exe.zip" "$cold_uri"
if command -v unzip >/dev/null; then
	unzip -d "$bin_dir" -o "$exe.zip"
else
	7z x -o"$bin_dir" -y "$exe.zip"
fi
chmod +x "$exe"
rm "$exe.zip"

echo "cold was installed successfully to $exe"
if command -v cold >/dev/null; then
	echo "Run 'cold --help' to get started"
else
	case $SHELL in
	/bin/zsh) shell_profile=".zshrc" ;;
	*) shell_profile=".bashrc" ;;
	esac
	echo "Manually add the directory to your \$HOME/$shell_profile (or similar)"
	echo "  export cold_INSTALL=\"$cold_install\""
	echo "  export PATH=\"\$cold_INSTALL/bin:\$PATH\""
	echo "Run '$exe --help' to get started"
fi
echo
echo "Stuck? Join our Discord https://discord.gg/cold"

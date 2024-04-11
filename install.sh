#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Loading animation function
spin() {
    spinner="/|\\-/|\\-"
    echo -ne " "
    while :
    do
        for i in {1..7}
        do
            echo -ne "\b${spinner:i:1}"
            sleep 0.1
        done
    done
}

# Permission check function
check_permissions() {
    if [[ $(id -u) -ne 0 ]]; then
        echo -e "${RED}Installer needs to move 'docmd' to /usr/local/bin, which requires permissions.${NC}"
        echo -e "You may be prompted for your password.${NC}"
        SUDO='sudo'
    else
        SUDO=''
    fi
}

# Determine OS
OS="unknown"
case "$(uname -s)" in
    Darwin)
        OS="darwin"
        ;;
    Linux)
        OS="linux"
        ;;
    *)
        echo -e "${RED}Unsupported OS.${NC}"
        exit 1
        ;;
esac

# Determine Architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64)
        ARCH="amd64"
        ;;
    arm64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}Unsupported architecture.${NC}"
        exit 1
        ;;
esac

# Define download URL with the provided GitHub link
URL="https://github.com/danecwalker/docmd/releases/latest/download/docmd-${OS}-${ARCH}"

# Check for permissions before proceeding
check_permissions

# Download the binary
echo -n "Downloading docmd for ${OS}-${ARCH}... "
spin &
SPIN_PID=$!
curl -L ${URL} -o docmd >/dev/null 2>&1 &
CURL_PID=$!

wait $CURL_PID
kill -9 $SPIN_PID
wait $SPIN_PID 2>/dev/null
echo -e "\n${GREEN}Download complete!${NC}"

# Make it executable
chmod +x docmd

# Move it to a directory in your PATH (e.g., /usr/local/bin)
echo -e "${GREEN}Installing docmd...${NC}"
${SUDO} mv docmd /usr/local/bin/docmd

echo -e "${GREEN}Installation complete. You can now use 'docmd'.${NC}"

#!/bin/bash
check_make() {
    if command -v make &> /dev/null; then
        echo "make already installed"
        return 0
    else
        echo "make not installed"
        return 1
    fi
}

install_make_macos() {
    echo "Installing make for macOS..."
    xcode-select --install
}

install_make_linux() {
    echo "Installing make for Linux..."

    if [ -x "$(command -v apt-get)" ]; then
        # Debian/Ubuntu
        sudo apt-get update
        sudo apt-get install -y make
    else
        echo "Package manager not compatible. Please install make manually"
        exit 1
    fi
}

if [[ "$OSTYPE" == "darwin"* ]]; then
    if ! check_make; then
        install_make_macos
    fi
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    if ! check_make; then
        install_make_linux
    fi
else
    echo "OS not compatible"
    exit 1
fi

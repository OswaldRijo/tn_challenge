#!/bin/sh

PROJECT=truenorthchallenge
BUF_VERSION=1.30.1

check_buf() {
    if command -v buf &> /dev/null; then
        echo "Buf already installed. Version: $(buf --version)"
        return 0
    else
        echo "Buf not installed"
        return 1
    fi
}

install_buf_linux() {
    echo "Installing Buf for Linux..."

    curl -sSL "https://github.com/bufbuild/buf/releases/download/v$BUF_VERSION/buf-Linux-x86_64.zip" -o "buf.zip"

    unzip buf.zip -d /usr/local/bin

    rm buf.zip
}

install_buf_macos() {
    echo "Installing Buf for macOS..."
    if ! command -v brew &> /dev/null; then
        echo "Installing Homebrew..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    fi
    curl -sSL "https://github.com/bufbuild/buf/releases/download/v$BUF_VERSION/buf-darwin-arm64.tar.gz" -o "buf.tar.gz"
    sudo tar -xzf buf.tar.gz
    sudo mv buf/bin/buf /usr/local/bin/
    sudo mv buf/bin/protoc-gen-buf-breaking /usr/local/bin/
    sudo mv buf/bin/protoc-gen-buf-lint /usr/local/bin/
    rm -rf buf
    rm buf.tar.gz
    buf --version
}

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"

if check_buf; then
    echo "Buf is ready to go."
else
    if [ "$OS" == "linux" ]; then
        install_buf_linux
    elif [ "$OS" == "darwin" ]; then
        install_buf_macos
    else
        echo "OS NOT SUPPORTED. Please, install Buf manually."
        exit 1
    fi

    if check_buf; then
        echo "Buf correctly installed."
    else
        echo "Error trying to install Buf."
        exit 1
    fi
fi

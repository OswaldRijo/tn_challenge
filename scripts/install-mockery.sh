#!/bin/bash

check_mockery() {
    if command -v mockery &> /dev/null; then
        echo "Mockery already installed"
        return 0
    else
        echo "Mockery not present."
        return 1
    fi
}

install_mockery_macos() {
    echo "Installing Mockery for macOS..."
    brew update
    brew install mockery
}

install_mockery_linux() {
    echo "Installing Mockery for Linux..."
    go env -w GO111MODULE=on
    go install github.com/vektra/mockery/v2@v2.43.0
}

if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    if ! check_mockery; then
        install_mockery_macos
    fi
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    if ! check_mockery; then
        install_mockery_linux
    fi
else
    echo "OS not compatible."
    exit 1
fi

if command -v mockery &> /dev/null; then
    echo "Mockery has been installed successfully."
else
    echo "There was an error installing Mockery."
    exit 1
fi

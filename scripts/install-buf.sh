#!/bin/sh

PROJECT=truenorthchallenge
BUF_VERSION=1.30.1
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"

check_buf() {
  n="$(uname -s | tr '[:upper:]' '[:lower:]')"

  if [ "$n" = "linux" ]; then
    return 1
  fi
  if command -v buf &> /dev/null; then
     buf_output=$(buf --version 2>&1)
     buf_status=$?

     if [ $buf_status -eq 0 ]; then
       echo "Buf is installed. Version: $buf_output"
       return 0
     else
       echo "Buf command found, but it failed to execute properly."echo"Error output: $buf_output"return 1
     fi
  else
     echo "Buf is not installed."
     return 1
  fi
}

install_buf_linux() {
    echo "Installing Buf for Linux..."

    curl -sSL "https://github.com/bufbuild/buf/releases/download/v$BUF_VERSION/buf-Linux-x86_64.tar.gz" -o "buf.tar.gz"

    sudo tar -xzf buf.tar.gz
    sudo mv buf/bin/buf /usr/local/bin/

    rm buf.tar.gz
    cd src/go && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && cd ../..
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



if check_buf; then
    echo "Buf is ready to go."
else
    if [ "$OS" = "linux" ]; then
        install_buf_linux
    elif [ "$OS" = "darwin" ]; then
        install_buf_macos
    else
        echo "OS NOT SUPPORTED. Please, install Buf manually."
        exit 1
    fi

    export PATH="$PATH:$(go env GOPATH)/bin"
fi

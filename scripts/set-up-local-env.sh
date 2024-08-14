#!/bin/sh

if command -v pnpm &> /dev/null; then
    echo "pnpm is already present. Version: $(pnpm --version)"
else
    echo "pnpm is not present. Downloading and instaling pnpm"

    if ! command -v node &> /dev/null; then
        echo "Node.js is not installed. Please, install Node.js first."
        exit 1
    fi

    npm install -g pnpm

    if command -v pnpm &> /dev/null; then
        echo "pnpm has been installed successfully. Version: $(pnpm --version)"
    else
        echo "pnpm not installed, please handle its installation manually."
    fi
fi


if command -v go &> /dev/null; then
    echo "Go is already installed. Version: $(go version)"
else
    echo "Go is not present. Downloading it"
    sudo ./scripts/install-go.sh
    echo "Go has been installed correctly. Version: $(go version)"
fi

sudo ./scripts/install-buf.sh
sudo ./scripts/run-mocks.sh

sudo make protos_go
sudo make protos_npm

sudo ./scripts/install-docker-compose.sh
sudo ./scripts/install-psql.sh
sudo ./scripts/set-up-queue-and-topic.sh


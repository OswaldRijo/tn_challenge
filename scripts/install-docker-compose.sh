#!/bin/sh
if command -v docker-compose &> /dev/null; then
    echo "docker-compose is already installed. Version: $(docker-compose --version)"
else
    echo "docker-compose is not installed. Downloading and installing docker-compose."

    DOCKER_COMPOSE_VERSION="2.29.0"

    OS="$(uname -s)"
    ARCH="$(uname -m)"

    if [ "$ARCH" == "x86_64" ]; then
        ARCH="x86_64"
    elif [ "$ARCH" == "aarch64" ]; then
        ARCH="aarch64"
    fi

    sudo curl -L "https://github.com/docker/compose/releases/download/v${DOCKER_COMPOSE_VERSION}/docker-compose-${OS}-${ARCH}" -o /usr/local/bin/docker-compose

    sudo chmod +x /usr/local/bin/docker-compose

    if command -v docker-compose &> /dev/null; then
        echo "docker-compose installed successfully. Version: $(docker-compose --version)"
    else
        echo "there was an error installing docker-compose. Please install it manually"
        exit 1
    fi
fi

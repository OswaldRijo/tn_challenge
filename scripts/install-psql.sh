#!/bin/sh
if command -v psql &> /dev/null; then
    echo "psql already installed. Version: $(psql --version)"
else
    echo "psql is not installed. Downloading and installing docker-compose."

    OS="$(uname -s | tr '[:upper:]' '[:lower:]')"

    if [ "$OS" == "linux" ]; then
        sudo apt update
        sudo apt install -y postgresql-client

        sudo systemctl stop postgresql
        sudo systemctl disable postgresql

    elif [ "$OS" == "darwin" ]; then
        if ! command -v brew &> /dev/null; then
            echo "Homebrew is not installed. Installing it now."
            /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        fi
        brew install postgresql

        brew services stop postgresql
    else
        echo "Os not supported. Please, install psql manually."
        exit 1
    fi

    if command -v psql &> /dev/null; then
        echo "psql has been installed correctly. Version: $(psql --version)"
    else
        echo "There was an error installing psql."
        exit 1
    fi
fi

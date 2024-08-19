#!/bin/sh

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
COMPOSE=""
if [ "$OS" = "linux" ]; then
  COMPOSE="docker compose"
  sudo systemctl stop postgresql
else
  COMPOSE="docker-compose"
fi

$COMPOSE -f docker/database/docker-compose.yaml up -d
sleep 5
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="user"
DB_NAME="truenorth_db"
export PGPASSWORD="pass"

DB_EXIST=$(psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME';")

if [ "$DB_EXIST" == "1" ]; then
    echo "The db '$DB_NAME' already exists."
else
    echo "The db '$DB_NAME' does not exists. Creating it now."
    psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -c "CREATE DATABASE $DB_NAME;"
fi

DB_NAME="truenorth_operations_db"
DB_EXIST=$(psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME';")

if [ "$DB_EXIST" == "1" ]; then
    echo "The db '$DB_NAME' already exists."
else
    echo "The db '$DB_NAME' does not exists. Creating it now."
    psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -c "CREATE DATABASE $DB_NAME;"
fi
unset PGPASSWORD

export $(grep -v '^#' ./.properties | xargs)
$COMPOSE -f docker/operations_service/docker-compose.yaml up -d
$COMPOSE -f docker/users_service/docker-compose.yaml up -d
$COMPOSE -f docker/public_api/docker-compose.yaml up -d
$COMPOSE -f docker/web_client/docker-compose.yaml up -d
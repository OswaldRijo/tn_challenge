#!/bin/sh
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
COMPOSE=""
if [ "$OS" = "linux" ]; then
  COMPOSE="docker compose"
  sudo systemctl stop postgresql
else
  COMPOSE="docker-compose"
fi

$COMPOSE -f docker/operations_service/docker-compose.yaml down
$COMPOSE -f docker/users_service/docker-compose.yaml down
$COMPOSE -f docker/public_api/docker-compose.yaml down
$COMPOSE -f docker/web_client/docker-compose.yaml down
$COMPOSE -f docker/database/docker-compose.yaml down
docker rmi public_api_tn_public_api
docker rmi operations_service_users_service
docker rmi operations_service_operations_service
docker rmi users_service_users_service
docker rmi web_client_tn_public_api
docker network rm backend postgres_db

#!/bin/sh
docker-compose -f docker/operations_service/docker-compose.yaml down
docker-compose -f docker/users_service/docker-compose.yaml down
docker-compose -f docker/public_api/docker-compose.yaml down
docker-compose -f docker/database/docker-compose.yaml down
docker rmi public_api_tn_public_api
docker rmi operations_service_users_service
docker rmi users_service_users_service
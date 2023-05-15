#!/bin/bash

set -e

scp -o StrictHostKeyChecking=no -i "$SSH_KEY" "${CI_PROJECT_DIR}/ci/docker-compose.yaml" "${REMOTE_USER}@${REMOTE_HOST}:~/lavka/docker-compose.yaml"
scp -o StrictHostKeyChecking=no -i "$SSH_KEY" "${CI_PROJECT_DIR}/.env" "${REMOTE_USER}@${REMOTE_HOST}:~/lavka/.env"

ENV_FILE="lavka/.env"
ssh -o StrictHostKeyChecking=no -i "$SSH_KEY" "$REMOTE_USER@$REMOTE_HOST" \
      "echo ${REGISTRY_PASS} | docker login -u ${REGISTRY_USER} --password-stdin &&
      docker compose -f ./lavka/docker-compose.yaml --env-file ${ENV_FILE} down || echo 'no lavka-compose' &&
      docker compose -f ./lavka/docker-compose.yaml --env-file ${ENV_FILE} up -d"
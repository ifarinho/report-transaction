#!/bin/bash

if [ -f .env ]; then
  source .env
else
  echo ".env file not found"
  exit 1
fi

docker run --rm \
  -e RUN_MODE="$RUN_MODE" \
  -e DB_POSTGRES_HOST="$DB_POSTGRES_HOST" \
  -e DB_POSTGRES_USER="$DB_POSTGRES_USER" \
  -e DB_POSTGRES_PASSWORD="$DB_POSTGRES_PASSWORD" \
  -e DB_POSTGRES_NAME="$DB_POSTGRES_NAME" \
  -e DB_POSTGRES_PORT="$DB_POSTGRES_PORT" \
  -e DB_POSTGRES_SSL_MODE="$DB_POSTGRES_SSL_MODE" \
  -e DB_POSTGRES_TIME_ZONE="$DB_POSTGRES_TIME_ZONE" \
  -e DB_POSTGRES_DATA_SOURCE_NAME="$DB_POSTGRES_DATA_SOURCE_NAME" \
  -e AWS_ACCESS_KEY_ID="$AWS_ACCESS_KEY_ID" \
  -e AWS_ACCESS_SECRET_KEY="$AWS_ACCESS_SECRET_KEY" \
  -e AWS_CREDENTIAL_TOKEN="$AWS_CREDENTIAL_TOKEN" \
  -e AWS_REGION="$AWS_REGION" \
  -e AWS_S3_BUCKET="$AWS_S3_BUCKET" \
  -e AWS_S3_PREFIX="$AWS_S3_PREFIX" \
  -e SERVICE_EMAIL="$SERVICE_EMAIL" \
  -e CORS_ORIGIN="$CORS_ORIGIN" \
  --network="host" \
  "$DOCKER_IMAGE" --filename "$1" --account "$2"

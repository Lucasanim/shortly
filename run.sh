#!/bin/bash
set -e 

docker compose up -d

echo "Waiting for DynamoDB to be ready..."
until nc -z localhost 8000; do
  sleep 1
done

echo "DynamoDB is up! Starting Go app..."

echo "Waiting for Redis be ready..."
until nc -z localhost 6379; do
  sleep 1
done

echo "Redis is up! Starting Go app..."
go run cmd/main.go

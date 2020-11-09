#!/bin/bash
docker build -t ironcore864/go-hello-http .
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push ironcore864/go-hello-http

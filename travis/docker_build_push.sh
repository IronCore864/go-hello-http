#!/bin/bash
VERSION=$(cat version)
docker build -t $DOCKER_REPO:$VERSION .
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push $DOCKER_REPO:$VERSION

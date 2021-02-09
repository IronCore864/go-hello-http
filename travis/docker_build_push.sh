#!/bin/bash
VERSION=$(cat version)

aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

docker build -t public.ecr.aws/u7w2h8y6/go-hello-http:$VERSION .

docker push public.ecr.aws/u7w2h8y6/go-hello-http:$VERSION

#!/bin/bash

# shellcheck disable=SC2164
cd ~/Fuever
git pull

docker-compose down
docker rmi avaqua/fuever:latest
docker build . -t avaqua/fuever:latest
docker-compose up -d

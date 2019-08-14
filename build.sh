#!/bin/bash

docker rmi jenkins:v1.0
docker build . -t jenkins:v1.0

docker rm -f jenkins
docker run -itd --restart=unless-stopped -v /etc/localtime:/etc/localtime -v /etc/timezone:/etc/timezone --name jenkins -v $(pwd):/data --network=host jenkins:v1.0

docker logs -f jenkins






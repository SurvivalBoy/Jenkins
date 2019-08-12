#!/bin/bash

docker rmi jenkins:v1.0
docker build . -t jenkins:v1.0

echo ./run.sh


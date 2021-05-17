#!/bin/bash
export TAG=$(git ls-remote --tags | awk '{print $2}' | grep -v '{}' | awk -F"/" '{print $3}' | tail -n 1)
export COMMIT=$(git ls-remote | grep refs/heads/main | cut -f 1)
echo $DOCKER_PASSWORD | docker login --username $DOCKER_USERNAME --password-stdin
docker build -t devopsunicorn/buoy:$COMMIT .
docker run -d --name test devopsunicorn/buoy:$COMMIT
docker ps -a
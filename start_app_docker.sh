#!/bin/bash

PORT="8080"
LAYOUT_PATH="../../assets/layout/"

docker build -t mss-image .
docker run -it --rm --name mss \
    -p $PORT:$PORT \
    -e LAYOUT_PATH=$LAYOUT_PATH \
    mss-image -port $PORT
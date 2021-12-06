#!/bin/bash

PORT="8080"

export LAYOUT_PATH="./assets/layout/" 

go run ./cmd/mss/main.go -port $PORT
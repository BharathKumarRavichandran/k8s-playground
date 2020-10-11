#!/bin/sh

set -x

#tail -f /dev/null

echo "######## Fetching Go dependencies ##########"
./scripts/install_deps.sh

sleep 5

echo "################## Starting server ##################"
go run main.go

#!/bin/sh

set -e

cd $1 && \
echo "Building docker image for $1..." && \
docker build . -q --tag galo_$1

docker container rm -f galo_$1 1>/dev/null 2>/dev/null || true && \
docker run --name galo_$1 -it galo_$1 "/main"

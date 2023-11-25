#!/bin/bash

set -e
docker run --entrypoint bash -it -v $(pwd):/app gcc -c "
cd /app;
g++ file.cpp relation.cpp main.cpp;
./a.out;
rm -rf ./a.out;"
#!/bin/bash

set -e
docker run --entrypoint bash -it -v $(pwd):/app gcc -c "
cd /app;
g++ strings.cpp  file.cpp relation.cpp datasource.cpp main.cpp;
./a.out;
rm -rf ./a.out;"
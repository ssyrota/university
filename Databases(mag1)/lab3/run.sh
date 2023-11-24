#!/bin/bash

set -e
docker run --entrypoint bash -it -v $(pwd):/app gcc -c "g++ /app/main.cpp; /a.out"
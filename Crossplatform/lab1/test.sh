#!/bin/bash

gcc -o liblinear_system.so -fpic -shared ./linear_system.c;
export GOARCH=arm64; export CGO_ENABLED=1; go build -o ./main main.go;
echo "Running main...";
./main;

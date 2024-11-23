#!/bin/bash

# relevant only for darwin
export DYLD_LIBRARY_PATH=../c_library/bin:$DYLD_LIBRARY_PATH;
export GOARCH=arm64; export CGO_ENABLED=1; go build -o main .;

./main;
rm -rf main;
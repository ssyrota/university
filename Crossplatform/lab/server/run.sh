#!/bin/bash

export CGO_ENABLED=1; go build -o main .;

# relevant only for darwin
export DYLD_LIBRARY_PATH=../c_library/bin:$DYLD_LIBRARY_PATH;

./main;
rm -rf main;
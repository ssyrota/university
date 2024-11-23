#!/bin/bash

# relevant only for darwin
export DYLD_LIBRARY_PATH=../c_library/bin:$DYLD_LIBRARY_PATH;
export CGO_ENABLED=1; go test -v -cover ./...;

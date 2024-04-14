#!/bin/bash

mkdir ./proto
protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=./proto ./map.proto
npx tsc; node ./dist/index.js
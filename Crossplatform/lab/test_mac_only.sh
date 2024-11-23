#!/bin/bash

mkdir -p ./bin;

echo "----PLATFORMS----";
darwin_arm64="arm64-darwin";

echo "darwin_arm64: $darwin_arm64";
echo ""
echo ""

echo "-----COMPILING LIBRARY-----";

echo "Building for $darwin_arm64...";
gcc -o ./bin/${darwin_arm64}_liblinear_system.so -fpic -shared ./linear_system.c;

echo "----COMPILATION SUCCESSFUL----";


echo ""
echo "----EXECUTION FOR DARWIN ARM64----";
cp ./bin/${darwin_arm64}_liblinear_system.so ./liblinear_system.so;
export GOARCH=arm64; export CGO_ENABLED=1; go build -o ./bin/${darwin_arm64}_main main.go;
echo "Running main for darwin arm64...";
./bin/${darwin_arm64}_main;
rm ./liblinear_system.so;

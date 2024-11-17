#!/bin/bash

mkdir -p ./bin;

echo "----PLATFORMS----";
darwin_arm64="arm64-darwin";
linux_arm64="arm64-linux";
linux_amd64="amd64-linux";

echo "darwin_arm64: $darwin_arm64";
echo "linux_arm64: $linux_arm64";
echo "linux_amd64: $linux_amd64";
echo ""
echo ""

echo "-----COMPILING LIBRARY-----";

echo "Building for $darwin_arm64...";
gcc -o ./bin/${darwin_arm64}_liblinear_system.so -fpic -shared ./linear_system.c;

echo "Building for $linux_arm64 using Docker...";
docker run --platform linux/arm64 --rm -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp arm64v8/gcc:latest gcc -o ./bin/${linux_arm64}_liblinear_system.so -fpic -shared ./linear_system.c;

echo "Building for $linux_amd64 using Docker...";
docker run --platform linux/amd64 --rm -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp gcc:latest gcc -o ./bin/${linux_amd64}_liblinear_system.so -fpic -shared ./linear_system.c;

echo "----COMPILATION SUCCESSFUL----";


echo ""
echo "----EXECUTION FOR DARWIN ARM64----";
cp ./bin/${darwin_arm64}_liblinear_system.so ./liblinear_system.so;
export GOARCH=arm64; export CGO_ENABLED=1; go build -o ./bin/${darwin_arm64}_main main.go;
echo "Running main for darwin arm64...";
./bin/${darwin_arm64}_main;
rm ./liblinear_system.so;

echo ""
echo "----EXECUTION FOR LINUX ARM64----"
docker run --platform linux/arm64 --rm -it -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:1.22 bash -c "\
  cp ./bin/${linux_arm64}_liblinear_system.so ./liblinear_system.so && \
  export GOARCH=arm64 && \
  export CGO_ENABLED=1 && \
  export LD_LIBRARY_PATH=./bin:\$LD_LIBRARY_PATH && \
  go build -o ./bin/${linux_arm64}_main . && \
  ./bin/${linux_arm64}_main; \
  rm ./liblinear_system.so; \
  echo 'Done';
"

echo ""
echo "----EXECUTION FOR LINUX AMD64----"
docker run --platform linux/amd64 --rm -it -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:1.22 bash -c "\
  cp ./bin/${linux_amd64}_liblinear_system.so ./liblinear_system.so && \
  export GOARCH=amd64 && \
  export CGO_ENABLED=1 && \
  export LD_LIBRARY_PATH=./bin:\$LD_LIBRARY_PATH && \
  go build -o ./bin/${linux_amd64}_main . && \
  ./bin/${linux_amd64}_main; \
  rm ./liblinear_system.so; \
  echo 'Done';
"


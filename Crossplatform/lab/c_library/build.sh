#!/bin/bash
source ../detect_platform.sh;

mkdir -p ./bin;

platform=$1;
platform=$(detect_platform $platform);

if [ "$platform" == "darwin_arm64" ]; then
  echo "Building for $platform...";
  gcc -o ./bin/lib${platform}_linear_system.so -fpic -shared ./linear_system.c;
fi

if [ "$platform" == "linux_arm64" ]; then
  docker run --platform linux/arm64 --rm -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp arm64v8/gcc:latest gcc -o ./bin/lib${platform}_linear_system.so -fpic -shared ./linear_system.c;
fi

if [ "$platform" == "linux_amd64" ]; then
  docker run --platform linux/amd64 --rm -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp gcc:latest gcc -o ./bin/lib${platform}_linear_system.so -fpic -shared ./linear_system.c;
fi

echo "----COMPILATION SUCCESSFUL----";

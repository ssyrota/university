#!/bin/bash
source ./detect_platform.sh;

platform=$1;
if [ -z "$platform" ]; then
  echo "NOTE: To run on different platform: $0 <platform>, available platforms: darwin_arm64, linux_arm64, linux_amd64";
  echo "";
  echo "Defaulting to detected platform";
  platform=$(detect_platform);
fi

echo "Running on platform: $platform";

echo "Building library...";
cd ./c_library;
./build.sh $platform;
cd ..;

echo "Testing server...";
if [ "$platform" == "darwin_arm64" ]; then
  echo "----TESTING FOR DARWIN ARM64----"
  cd server;
  ./test.sh;
  cd ..;
fi

if [ "$platform" == "linux_arm64" ]; then
  echo "----TESTING FOR LINUX ARM64----"
  docker run --platform linux/arm64 --rm -it -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:1.22 bash -c "\
    export GOARCH=arm64;
    export CGO_ENABLED=1;
    cd ./server;
    export LD_LIBRARY_PATH=/usr/src/myapp/c_library/bin:\$LD_LIBRARY_PATH;
    ./test.sh;
    echo 'Done';
  "
fi

if [ "$platform" == "linux_amd64" ]; then
  echo "----TESTING FOR LINUX AMD64----"
  docker run --platform linux/amd64 --rm -it -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:1.22 bash -c "\
    export GOARCH=amd64;
    export CGO_ENABLED=1;
    cd ./server;
    export LD_LIBRARY_PATH=/usr/src/myapp/c_library/bin:\$LD_LIBRARY_PATH;
    ./test.sh;
    echo 'Done';
  "
fi
#!/bin/bash
source ./detect_platform.sh;

echo "Building library...";
cd ./c_library;
./build.sh $platform;
cd ..;

echo "Usage: $0 <platform> <port>";
echo "";

platform=$1;
if [ -z "$platform" ]; then
  echo "NOTE: To run on different platform: $0 <platform>, available platforms: darwin_arm64, linux_arm64, linux_amd64";
  echo "";
  echo "Defaulting to detected platform";
  platform=$(detect_platform);
  echo "Running application on platform: $platform";
fi

port=$2;
if [ -z "$port" ]; then
  port="8080";
  echo "Defaulting to port: $port";
fi

if [ "$platform" == "darwin_arm64" ]; then
  echo "----RUNNING ON DARWIN ARM64----"
  cd server;
  export PORT=$port;
  ./run.sh;
  cd ..;
fi

if [ "$platform" == "linux_arm64" ]; then
  echo "----RUNNING ON LINUX ARM64----"
  docker run -p $port:$port --platform linux/arm64 --rm -it -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:1.22 bash -c "\
    export GOARCH=arm64;
    export CGO_ENABLED=1;
    cd ./server;
    export PORT=$port;
    export LD_LIBRARY_PATH=/usr/src/myapp/c_library/bin:\$LD_LIBRARY_PATH;
    ./run.sh;
    echo 'Done';
  "
fi

if [ "$platform" == "linux_amd64" ]; then
  echo "----RUNNING ON LINUX AMD64----"
  docker run -p $port:$port --platform linux/amd64 --rm -it -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:1.22 bash -c "\
    export GOARCH=amd64;
    export PORT=$port;
    export CGO_ENABLED=1;
    cd ./server;
    export LD_LIBRARY_PATH=/usr/src/myapp/c_library/bin:\$LD_LIBRARY_PATH;
    ./run.sh;
    echo 'Done';
  "
fi
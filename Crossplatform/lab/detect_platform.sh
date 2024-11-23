#!/bin/bash

detect_platform(){
  platform=$1;

  if [ "$platform" != "" ]; then
    if [ "$platform" != "darwin_arm64" ] && [ "$platform" != "linux_arm64" ] && [ "$platform" != "linux_amd64" ]; then
      echo "Unsupported platform: $platform";
      exit 1;
    fi
  fi

  arch=$(uname -m);
  os=$(uname -s);
  if [ "$os" == "Darwin" ] && [ "$arch" == "arm64" ]; then
      platform="darwin_arm64";
  elif [ "$os" == "Linux" ] && [ "$arch" == "arm64" ]; then
    platform="linux_arm64";
  elif [ "$os" == "Linux" ] && [ "$arch" == "amd64" ]; then
    platform="linux_amd64";
  else
    echo "Unsupported platform: $os $arch";
    exit 1;
  fi

  echo $platform;
}

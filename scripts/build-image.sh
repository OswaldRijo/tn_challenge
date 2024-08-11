#!/usr/bin/env bash

if [ -d "./build" ]; then
  rm -r ./build
fi

if [ ! -d "./build" ]; then
  mkdir ./build
  mkdir ./build/src
  mkdir ./build/src/$BUILD_PATH
fi

cp -r "./scripts" ./build/
cp -r "./migrations" ./build/
cp -r "./docker" ./build/
cp -r "./src" ./build/
sleep 5
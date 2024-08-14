#!/usr/bin/env bash
cd src/node/pb && pnpm install --save-dev ts-protoc-gen && cd ../../..
cd src/go && go build -o gen packages/mock_generator/cmd/main.go
chmod +x gen
./gen && rm gen



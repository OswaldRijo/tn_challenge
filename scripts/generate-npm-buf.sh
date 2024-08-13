#!/bin/sh
PROJECT=true_north_challenge
# buf is installed to ~/bin/your-project-name.
PATH=$PATH:$GOPATH/bin
BIN_DIR=$HOME/bin/$PROJECT

cd src/node/pb && $BIN_DIR/buf generate ./../../../protobuf
pnpm run generate-index

echo "Command ran successfully"
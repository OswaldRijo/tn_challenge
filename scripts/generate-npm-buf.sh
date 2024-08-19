#!/bin/sh
PROJECT=tn
# buf is installed to ~/bin/your-project-name.
PATH=$PATH:$GOPATH/bin
BIN_DIR=$HOME/bin/$PROJECT

cd npm && $BIN_DIR/buf generate ./../protobuf
npm run generate-index
npm run build

echo "Command ran successfully"
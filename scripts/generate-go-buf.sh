#!/bin/sh
PROJECT=tn
# buf is installed to ~/bin/your-project-name.
PATH=$PATH:$GOPATH/bin
BIN_DIR=$HOME/bin/$PROJECT

$BIN_DIR/buf generate ./protobuf
go build -o gen src/generator/cmd/main.go
chmod +x gen
./gen && rm gen && rm -rf $HOME/bin
#!/bin/sh

PROJECT=truenorthchallenge
# Use your desired buf version
BUF_VERSION=1.30.1
# buf is installed to ~/bin/your-project-name.
PATH=$PATH:$GOPATH/bin
BIN_DIR=$HOME/bin/$PROJECT
mkdir -p $BIN_DIR
mkdir -p "$HOME/go/pb"

curl -sSL \
	"https://github.com/bufbuild/buf/releases/download/v$BUF_VERSION/buf-$(uname -s)-$(uname -m)" \
	-o "$BIN_DIR/buf"
chmod +x "$BIN_DIR/buf"

echo $GOPATH
echo $GOBIN

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

UNAME = $(shell uname -s)
LOWER_UNAME = $(shell echo $(UNAME) | tr A-Z a-z)
PROTO_DIR_GEN_GO = src/go/pb
PROTO_DIR_GEN_NODE = src/node/pb


ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
	OS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
	CHECK_DIR_CMD = if (!(Test-Path $@)) { $$e = [char]27; Write-Error "$$e[31mDirectory $@ doesn't exist$${e}[0m" }
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
	RM_F_CMD = Remove-Item -erroraction sil entlycontinue -Force
	RM_RF_CMD = ${RM_F_CMD} -Recurse
	SERVER_BIN = ${SERVER_DIR}.exe
else
	SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS = macos ${VERSION_AND_ARCH}
	else
		OS = linux ${VERSION_AND_ARCH}
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
	SERVER_BIN = ${SERVER_DIR}
endif

run_go_fmt:
	cd src/go && go fmt;

run_go_generate:
	cd src/go && go generate ./...;

buf:
	buf generate protobuf
	@echo "Command ran successfully";

go_mocks:
	@echo "Running go mocks"
	cd src/go && go generate ./...

test: protos ## Launch tests
	go test ./...

coverage: protos ## Launch tests
	go test ./... -cover

rebuild: clean protos ## Rebuild the whole project

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@${HELP_CMD}

clean_protos:
	rm -rf ${PROTO_DIR_GEN_NODE}/src
	mkdir ${PROTO_DIR_GEN_NODE}/src
	touch ${PROTO_DIR_GEN_NODE}/src/index.ts
	rm -rf ${PROTO_DIR_GEN_GO}

protos: protos_go protos_npm

protos_go: clean_protos buf run_pb_mocks run_go_generate run_go_fmt

protos_npm: buf_node node_generate_index pb_link_dir

go_generate:
	cd src/go && go generate $(dir)

buf_node:
	cd ${PROTO_DIR_GEN_NODE} && buf generate ./../../../protobuf
	@echo "Command ran successfully";

pb_link_dir:
	rm -rf src/node/public_api/src/pb
	cp -r ${PROTO_DIR_GEN_NODE}/src src/node/public_api/src/pb
	@echo "Command ran successfully";

node_generate_index:
	pnpm run --prefix src/node/pb generate-index
	@echo "Index.ts created successfully";

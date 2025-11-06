.PHONY: default lint test race run linux darwin macos windows linux32 darwin32 macos32 windows32 all all32 release

name?=$(shell basename `pwd`)
ProjectName=$(name)
ProjectDir=$(shell pwd)
OSName=`uname -s | tr 'A-Z' 'a-z'`
OSARCH:=
ifeq ($(OS),Windows_NT)
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		OSARCH = amd64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),x86)
		OSARCH = 386
	endif
else
	UNAME_P := $(shell uname -m)
	ifeq ($(UNAME_P),x86_64)
		OSARCH = amd64
	endif
	ifneq ($(filter %86,$(UNAME_P)),)
		OSARCH = 386
	endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		OSARCH = arm
	endif
endif
OutputDir=$(ProjectDir)/_bin/$(ProjectName)
MainFilePath=$(ProjectDir)/cmd/$(ProjectName)
GoModName=`head -n 1 go.mod | cut -d ' ' -f 2`
PkgList=$(shell go list ./... | grep -v /vendor/)
BranchName=`git symbolic-ref --short -q HEAD`
ifeq ($(FOUND_BOOK),)
	ifneq ($(CI_COMMIT_REF_NAME),)
		BranchName=$(CI_COMMIT_REF_NAME)
	endif
endif
BuildGoVersion=`go version|sed -e 's/go version //g'`
BuildTime=`date '+%Y%m%d.%H%M%S'`
BuildFlags=-ldflags "-extldflags=-static -X '$(GoModName)/internal/pkg/manager.GITBRANCH=$(BranchName)' -X $(GoModName)/internal/pkg/manager.GITHASH=`git rev-parse HEAD | cut -c1-8` -X $(GoModName)/internal/pkg/manager.BUILDTIME=$(BuildTime) -X '$(GoModName)/internal/pkg/manager.GOVERSION=$(BuildGoVersion)'"
ifneq ($(tags),)
	BuildFlags:=$(BuildFlags) -tags "$(tags)"
endif

default:
	@go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-$(OSName)-$(OSARCH) $(MainFilePath)

lint:
	@golangci-lint run --exclude-use-default=false ./...

test:
	@go test -v $(PkgList)

race:
	@CGO_ENABLED=1 go test -race -short $(PkgList)

run:
	@go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-$(OSName)-$(OSARCH) $(MainFilePath)
	@cd $(OutputDir) && GO_DEBUG=true ./$(ProjectName)-$(OSName)-$(OSARCH)

all: linux darwin windows
all32: linux32 darwin32 windows32

linux:
	@GOOS=linux GOARCH=amd64 go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-linux-amd64 $(MainFilePath)

darwin macos:
	@GOOS=darwin GOARCH=amd64 go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-darwin-amd64 $(MainFilePath)

windows:
	@GOOS=windows GOARCH=amd64 go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-windows-amd64 $(MainFilePath)

linux32:
	@GOOS=linux GOARCH=386 go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-linux-386 $(MainFilePath)

darwin32 macos32:
	@GOOS=darwin GOARCH=386 go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-darwin-386 $(MainFilePath)

windows32:
	@GOOS=windows GOARCH=386 go build $(BuildFlags) -o $(OutputDir)/$(ProjectName)-windows-386 $(MainFilePath)

release: # for gitlab-ci
	@GOOS=linux GOARCH=amd64 go build $(BuildFlags) -o $(ProjectName)-linux-amd64 $(MainFilePath)

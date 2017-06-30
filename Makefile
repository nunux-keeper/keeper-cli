.SILENT :
.PHONY : clean build

# Username
USERNAME=nunux-keeper

# App name
APPNAME=keepctl

# Base image
BASEIMAGE=golang:1.8

# Go configuration
GOOS?=linux
GOARCH?=amd64

# Extract version infos
VERSION:=`git describe --tags`
LDFLAGS=-ldflags "-X github.com/nunux-keeper/keeper-cli/version.App=${VERSION}"

all: clean build

# Include common Make tasks
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
makefiles:=$(root_dir)/makefiles
include $(makefiles)/help.Makefile
include $(makefiles)/docker.Makefile

glide.lock:
	glide install

## Clean built files
clean:
	-rm -rf release

## Build executable
build: glide.lock
	mkdir -p release
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o release/$(APPNAME)-$(GOOS)-$(GOARCH)


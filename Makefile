.SILENT :

# App name
APPNAME=keepctl

# Base image
BASEIMAGE=golang:1.8

# Go configuration
GOOS?=linux
GOARCH?=amd64

# Add exe extension if windows target
is_windows:=$(filter windows,$(GOOS))
EXT:=$(if $(is_windows),".exe","")

# Go app path
APP_BASE=${GOPATH}/src/github.com/nunux-keeper

# Artefact name
ARTEFACT=release/$(APPNAME)-$(GOOS)-$(GOARCH)$(EXT)

# Extract version infos
VERSION:=`git describe --tags`
LDFLAGS=-ldflags "-X github.com/nunux-keeper/keeper-cli/version.App=${VERSION}"

all: build

# Include common Make tasks
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
makefiles:=$(root_dir)/makefiles
include $(makefiles)/help.Makefile
include $(makefiles)/docker.Makefile

$(APP_BASE)/keeper-cli:
	echo "Creating GO src link: $(APP_BASE)/keeper-cli ..."
	mkdir -p $(APP_BASE)
	ln -s $(root_dir) $(APP_BASE)/keeper-cli

glide.lock:
	echo "Installing dependencies ..."
	glide install

## Clean built files
clean:
	-rm -rf release
.PHONY : clean

## Build executable
build: glide.lock $(APP_BASE)/keeper-cli
	mkdir -p release
	echo "Building: $(ARTEFACT) ..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(ARTEFACT)
.PHONY : build

$(ARTEFACT): build

## Install executable
install: $(ARTEFACT)
	echo "Installing release/$(APPNAME)-$(GOOS)-$(GOARCH)$(EXT) to ${HOME}/.local/bin/keepctl ..."
	cp release/$(APPNAME)-$(GOOS)-$(GOARCH)$(EXT) ${HOME}/.local/bin/keepctl
.PHONY : install


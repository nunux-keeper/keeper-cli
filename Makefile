package = github.com/nunux-keeper/keeper-cli
name = keepctl

.PHONY: release

VERSION:=`git describe --tags`
LDFLAGS=-ldflags "-X github.com/ncarlier/keeper-cli/version.App=${VERSION}"

release:
	go get -v github.com/spf13/cobra/cobra
	go get -v github.com/bgentry/speakeasy
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o release/$(name)-linux-amd64 $(package)
	GOOS=linux GOARCH=arm go build $(LDFLAGS) -o release/$(name)-linux-arm $(package)

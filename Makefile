package = github.com/ncarlier/keeper-cli
name = keepctl

.PHONY: release

release:
	go get -v github.com/spf13/cobra/cobra
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(name)-linux-amd64 $(package)
	GOOS=linux GOARCH=arm go build -o release/$(name)-linux-arm $(package)

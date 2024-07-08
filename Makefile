.DEFAULT_GOAL := dev
NAME = playgo
OS = linux
ARCH = amd64
VERSION := $(shell git describe --tags --abbrev=0)

build:
	@GOOS=linux GOARCH=amd64 go build -o bin/${NAME}-amd64-linux-${VERSION} -ldflags "-X github.com/axyut/playgo/cmd.Version=${VERSION}" .
	@GOOS=darwin GOARCH=amd64 go build -o bin/${NAME}-amd64-darwin-${VERSION} -ldflags "-X github.com/axyut/playgo/cmd.Version=${VERSION}" .
	@GOOS=windows GOARCH=amd64 go build -o bin/${NAME}-amd64-windows-${VERSION} -ldflags "-X github.com/axyut/playgo/cmd.Version=${VERSION}" .

build_arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/${NAME}-arm64-darwin-${VERSION} -ldflags "-X github.com/axyut/playgo/cmd.Version=${VERSION}" .
	GOOS=windows GOARCH=arm64 go build -o bin/${NAME}-arm64-windows-${VERSION} -ldflags "-X github.com/axyut/playgo/cmd.Version=${VERSION}" .
	GOOS=linux GOARCH=arm64 go build -o bin/${NAME}-arm64-linux-${VERSION} -ldflags "-X github.com/axyut/playgo/cmd.Version=${VERSION}" .
# https://issueantenna.com/repo/ebitengine/oto/issues/238
# https://stackoverflow.com/questions/51475992/cgo-cross-compiling-from-amd64linux-to-arm64linux
clean:
	rm -rf bin

test:
	@go test -v ./...
	
dev:
	@go run . 

release:
	@goreleaser release --snapshot --clean

run: build
	@bin/${NAME}-${ARCH}-${OS}-${VERSION}

install:
	sudo apt install libasound2-dev
	go mod tidy

lint:
	@golangci-lint run

git: lint test # make git m="your message"
	git add .
	git commit -m "$(m)"
	git push

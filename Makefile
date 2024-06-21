.DEFAULT_GOAL := dev
NAME = playgo
OS = linux
ARCH = amd64
VERSION = 0.1.3

build:
	@GOOS=linux GOARCH=amd64 go build -o bin/${NAME}-amd64-linux-${VERSION} .
	@GOOS=darwin GOARCH=amd64 go build -o bin/${NAME}-amd64-darwin-${VERSION} .
	@GOOS=windows GOARCH=amd64 go build -o bin/${NAME}-amd64-windows-${VERSION} .

build_arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/${NAME}-arm64-darwin-${VERSION} .
	GOOS=windows GOARCH=arm64 go build -o bin/${NAME}-arm64-windows-${VERSION} .
	GOOS=linux GOARCH=arm64 go build -o bin/${NAME}-arm64-linux-${VERSION} .

clean:
	rm -rf bin

test:
	@go test -v ./...
	
dev:
	@go run . .

run: build
	@bin/${NAME}-${ARCH}-${OS}-${VERSION}

install:
	@go mod tidy

lint:
	golangci-lint run



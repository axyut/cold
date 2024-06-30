.DEFAULT_GOAL := dev
NAME = playgo
OS = linux
ARCH = amd64

build:
	@GOOS=linux GOARCH=amd64 go build -o bin/${NAME}-amd64-linux .
	@GOOS=darwin GOARCH=amd64 go build -o bin/${NAME}-amd64-darwin .
	@GOOS=windows GOARCH=amd64 go build -o bin/${NAME}-amd64-windows .

build_arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/${NAME}-arm64-darwin .
	GOOS=windows GOARCH=arm64 go build -o bin/${NAME}-arm64-windows .
	GOOS=linux GOARCH=arm64 go build -o bin/${NAME}-arm64-linux .

clean:
	rm -rf bin

test:
	@go test -v ./...
	
dev:
	@go run . 

run: build
	@bin/${NAME}-${ARCH}-${OS}

install:
	@go mod tidy

lint:
	golangci-lint run

# all: build build_arm

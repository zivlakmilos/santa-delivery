.PHONY: build test

all: run

run: build
	./build/santadelivery

build:
	@GOOS=linux go build -o build/santadelivery ./cmd/santadelivery

test:
	go test ./...


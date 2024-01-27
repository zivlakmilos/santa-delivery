.PHONY: build build-editor test

all: run

run: build
	./build/santadelivery

build:
	@GOOS=linux go build -o build/santadelivery ./cmd/santadelivery

run-editor: build-editor
	./build/leveleditor

build-editor:
	@GOOS=linux go build -o build/leveleditor ./cmd/leveleditor

test:
	go test ./...


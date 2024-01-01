.PHONY: build

all: run

run: build
	./build/santadelivery

build:
	@GOOS=linux go build -tags prod -o build/santadelivery ./cmd/santadelivery

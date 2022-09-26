.PHONY: build

FILENAME=main

all: build
build:
	# @CGO_ENABLED=0 GOOS=android GOARCH=arm64 go build -ldflags="-s -w" -o "build/${FILENAME}"
	@CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o "build/${FILENAME}"
	@echo "finish <3"




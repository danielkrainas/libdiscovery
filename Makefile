SRC_PKGS=$(shell go list ./... | grep -v vendor)

.PHONY: clean test

all: compile

clean:
	go clean ./...

compile:
	go build .
	go build ./locator/...

test:
	set -e; 
	for pkg in $(SRC_PKGS); \
	do \
		go test -v $$pkg; \
	done 

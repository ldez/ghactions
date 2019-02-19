.PHONY: clean fmt check test

GOFILES := $(shell git ls-files '*.go' | grep -v '^vendor/')

default: clean check test

clean:
	rm -rf cover.out

test: clean
	GO111MODULE=on go test -v -cover ./...

check:
	GO111MODULE=on golangci-lint run

fmt:
	gofmt -s -l -w $(GOFILES)

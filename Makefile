.PHONY: clean fmt check test

default: clean check test

clean:
	rm -rf cover.out

test: clean
	GO111MODULE=on go test -v -cover ./...

check:
	GO111MODULE=on golangci-lint run

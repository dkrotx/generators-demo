.PHONY: go-generate
default: run

bin/gowrap:
	go build -mod=readonly -o bin/gowrap github.com/hexdigest/gowrap/cmd/gowrap

go-generate: bin/gowrap
	PATH=$$PWD/bin:$$PATH go generate ./...

run:
	$(MAKE) go-generate
	go run .

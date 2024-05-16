VERSION := $(shell git describe --abbrev=0 --tags)
.PHONY: clean test gen

build: gen aws-sdk-client-go
	

aws-sdk-client-go: go.* *.go
	CGO_ENABLED=0 \
		go build -o $@ \
		-tags netgo \
		-ldflags '-s -w -extldflags "-static" -X github.com/fujiwara/aws-sdk-client-go.Version=$(VERSION)' \
		cmd/aws-sdk-client-go/main.go

clean:
	rm -rf *_gen.go aws-sdk-client-go dist/ cmd/aws-sdk-client-gen/gen.go
	go mod tidy

gen:
	go generate ./cmd/aws-sdk-client-gen .
	go fmt

test:
	go test -v .

packages:
	goreleaser build --skip=validate --clean

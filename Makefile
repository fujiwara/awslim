VERSION := $(shell git describe --tags)
.PHONY: clean test gen

build: gen awslim
	

awslim: go.* *.go
	CGO_ENABLED=0 \
		go build -o $@ \
		-tags netgo \
		-ldflags '-s -w -extldflags "-static" -X github.com/fujiwara/awslim.Version=$(VERSION)' \
		cmd/awslim/main.go

clean:
	rm -rf *_gen.go awslim dist/ cmd/awslim-gen/gen.go
	go mod tidy

gen:
	go generate ./cmd/awslim-gen .
	go fmt

test:
	go test -v .

packages:
	goreleaser build --skip=validate --clean

docker-build-and-push: Dockerfile
	docker buildx build \
	--platform=linux/amd64,linux/arm64 \
	-t ghcr.io/fujiwara/awslim:builder \
	-f Dockerfile \
	--push \
	.

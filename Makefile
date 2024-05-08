.PHONY: clean test

aws-sdk-client-go: go.* *.go gen
	go build -o $@ cmd/aws-sdk-client-go/main.go

clean:
	rm -rf *_gen.go aws-sdk-client-go dist/ cmd/aws-sdk-client-gen/gen.go
	go mod tidy

gen:
	go generate ./cmd/aws-sdk-client-gen .
	go fmt

test:
	go test -v ./...

install:
	go install github.com/fujiwara/aws-sdk-client-go/cmd/aws-sdk-client-go

dist:
	goreleaser build --snapshot --rm-dist

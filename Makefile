.PHONY: clean test

aws-sdk-client-go: go.* *.go gen
	go build -o $@ cmd/aws-sdk-client-go/main.go

gen:
	go get ./...
	go run cmd/aws-sdk-client-gen/main.go
	go fmt .

clean:
	rm -rf *_gen.go aws-sdk-client-go dist/

test:
	go test -v ./...

install:
	go install github.com/fujiwara/aws-sdk-client-go/cmd/aws-sdk-client-go

dist:
	goreleaser build --snapshot --rm-dist

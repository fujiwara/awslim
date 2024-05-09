.PHONY: clean test gen

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

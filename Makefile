.PHONY: clean test gen

aws-sdk-client-go: go.* *.go gen
	CGO_ENABLED=0 \
		go build -o $@ \
		-tags netgo \
		-ldflags '-s -w -extldflags "-static"' \
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
	goreleaser build --skip-validate --clean

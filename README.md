# aws-sdk-client-go

aws-sdk-client-go is a Go client for AWS SDK services.

This client is auto generated from the AWS SDK Go service client.

## Build

aws-sdk-client-go does not release binaries. Because the binary containing all services is too large, you need to build the client yourself to support the services you want to use.

The client is built by a configuration file `gen.yaml`.

```yaml
services:
  ecs:
    - DescribeClusters
    - DescribeTasks
  firehose:
    - DescribeDeliveryStream
    - ListDeliveryStreams
  kinesis:
    # all methods of the service
```

Keys of `services` are AWS service names (`github.com/aws/aws-sdk-go-v2/service/*`), and values are method names of the service client. If you don't specify the method names, all methods of the service client are generated.

To build the client, run the following commands:

```console
$ go generate ./cmd/aws-sdk-client-gen .
$ go build -o your-client ./cmd/aws-sdk-client-go/main.go
```

1. `go generate ./cmd/aws-sdk-client-gen .` generates the generator by `gen.yaml`.
2. `go build -o your-client ./cmd/aws-sdk-client-go/main.go` builds your client.

## Usage

```console
$ aws-sdk-client-go [service] [method] [input]
```

- `service`: AWS service name.
- `method`: Method name of the service client.
- `input`: JSON input for the method.

The output is JSON format.

### Examples

#### Show supported services

```console
$ aws-sdk-client-go
```

#### Show methods of the service

```console
$ aws-sdk-client-go ecs
```

#### Call method of the service

```console
$ aws-sdk-client-go ecs DescribeClusters '{"Cluster":"default"}'
```

The third argument is JSON input for the method. If the method does not require input, you can omit the third argument (implicitly `{}` passed).


## LICENSE

MIT

## Author

Fujiwara Shunichiro

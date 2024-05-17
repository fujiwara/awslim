# aws-sdk-client-go

aws-sdk-client-go is a Go client CLI for AWS services.

This CLI is auto generated from the AWS SDK Go v2 service client.

## Motivation

The [AWS CLI](https://aws.amazon.com/cli/) is very useful, but it requires too many CPU and memory resources to boot up. This client is a simplified alternative to the AWS CLI for limited use cases.

## Install

**Note:** The release binaries are large (about 500MB after being extracted) and slow a bit to boot up (about 100ms), because they contain codes for access to **[all AWS services](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service)**. If you want to use only a few services, you can build the client yourself. See [Build](#build) section.

### Release binary

Download the binary from the [release page](https://github.com/fujiwara/aws-sdk-client-go/releases).

### Homebrew

```console
$ brew install fujiwara/tap/aws-sdk-client-go
```

## Build

You can build the client yourself, including only the needed services and methods. The optimized binary is small and boots up quickly.

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

Keys of `services` are AWS service names (`github.com/aws/aws-sdk-go-v2/service/*`), and values are method names of the service client (for example, `s3` is [s3.Client](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#Client)). If you don't specify the method names, all methods of the service client are generated.

To build the client, run the following commands (or simply run `make`):

```console
$ go generate ./cmd/aws-sdk-client-gen .
$ go build -o your-client ./cmd/aws-sdk-client-go/main.go
```

1. `go generate ./cmd/aws-sdk-client-gen .` generates the generator by `gen.yaml`.
2. `go build -o your-client ./cmd/aws-sdk-client-go/main.go` builds your client.


## Performance comparison

Example of execution `sts get-caller-identity` on 0.25vCPU Fargate(AMD64).
`/usr/bin/time -v` is used for measurement.

| command | CPU time(user, sys)| Elapsed time(s) | Max memory(MB) |
| ---- | ---- | ---- | --- |
| aws               | 0.67 + 0.10 = 0.77 | 3.11 | 64.2 |
| aws-sdk-client-go(all) | 0.08 + 0.03 = 0.11 | 0.43 | 101.5 |
| aws-sdk-client-go(40) | 0.02 + 0.01 = 0.03 | 0.05 | 30.2 |

- `aws-sdk-client-go`(built for all AWS services): 7.0x faster than `aws`
- `aws-sdk-client-go`(built for 40 AWS services): 62.0x faster than `aws`

`aws-cli/2.15.51 Python/3.11.8`, `aws-sdk-client-go 0.0.10`

## Usage

```
Usage: aws-sdk-client-go [<service> [<method> [<input>]]] [flags]

Arguments:
  [<service>]    service name
  [<method>]     method name
  [<input>]      input JSON

Flags:
  -h, --help            Show context-sensitive help.
  -c, --compact         compact JSON output
  -q, --query=STRING    JMESPath query to apply to output
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

If the method name is "kebab-case", it automatically converts to "PascalCase" (for example, `describe-clusters` -> `DescribeClusters`).

```console
$ aws-sdk-client-go ecs describe-clusters '{"Cluster":"default"}'
```

#### Query output by JMESPath

`--query` option allows you to query the output by JMESPath like the AWS CLI.

```console
$ aws-sdk-client-go ecs DescribeClusters '{"Cluster":"default"}' \
  --query 'Clusters[0].ClusterArn'
```

#### Show help

aws-sdk-client-go is a simple wrapper of the AWS SDK Go v2 service client. Its usage is the same as that of the AWS SDK Go v2. If the third argument is "help", it shows the URL of the method documentation.

```console
$ aws-sdk-client-go ecs DescribeClusters help
See https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/ecs#Client.DescribeClusters
```

## LICENSE

MIT

## Author

Fujiwara Shunichiro

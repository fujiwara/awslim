# aws-sdk-client-go

`aws-sdk-client-go` is a CLI for AWS services by [Go](https://go.dev/). This CLI is generated from the [AWS SDK Go v2](https://github.com/aws/aws-sdk-go-v2) service client.

## Motivation

While the [AWS CLI](https://aws.amazon.com/cli/) is very usefull, it can be resource intensive to boot up. `aws-sdk-client-go` offers a simpler and faster alternative for limited use cases. It acts as a simple wrapper around AWS SDK Go v2, providing essential functionality without the full feature set of the AWS CLI.


### Features

- Call any method(API) of the AWS service client.
- Use JSON or Jsonnet for input.
- Output the result in JSON format.
- Bind a file to the input/output struct of the method.
- Query the output by JMESPath.
- Use the AWS CLI configuration file. (i.e., `~/.aws/config`)

### Limitations

- Not 100% compatible with the AWS CLI.
- No support for AWS CLI plugins. (i.e., `session-manager-plugin`)

## Installation

**Note:** The release binaries are large (about 500MB after being extracted) and have a slight startup delay (about 100ms), due to the inclusion of code to access **[all AWS services](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service)**. However, it is still [faster](#performance-comparison) than the AWS CLI.

For optimized performance, build your own binary tailored to the specific services you need. See the [Build](#build) section for details.

### Release binary

Download the binary from the [release page](https://github.com/fujiwara/aws-sdk-client-go/releases).

### Homebrew

```console
$ brew install fujiwara/tap/aws-sdk-client-go
```

## Build

You can build the client yourself, including only the needed services and methods. This produces a smaller, faster binary.

The client is built by using a `AWS_SDK_CLIENT_GO_GEN` environment variable or a `gen.yaml` configuration file.

### `AWS_SDK_CLIENT_GO_GEN` environment variable

Set the `AWS_SDK_CLIENT_GO_GEN` environment variable to specify the services you want to include, separated by commas.

For example, to build the client for ECS, Firehose, and S3:

```console
$ export AWS_SDK_CLIENT_GO_GEN="ecs,firehose,s3"
```

This will generate all methods for the specified services. If you want to generate only specific methods, use the gen.yaml configuration file.

### `gen.yaml` configuration file

```yaml
# gen.yaml
services:
  ecs:
    - DescribeClusters
    - DescribeTasks
  firehose:
    - DescribeDeliveryStream
    - ListDeliveryStreams
  s3:
    # all methods of the service
```

Keys under `services` are AWS service names (`github.com/aws/aws-sdk-go-v2/service/*`), and values are method names of the service client (for example, `s3` is [s3.Client](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#Client)). If you don't specify the method names, all methods of the service client are generated.

### Build on your machine

To build the client, run the following commands (or simply run `make`):

```console
$ go generate ./cmd/aws-sdk-client-gen .
$ go build -o your-client ./cmd/aws-sdk-client-go/main.go
```

1. `go generate ./cmd/aws-sdk-client-gen .` generates the generator by `gen.yaml`.
2. `go build -o your-client ./cmd/aws-sdk-client-go/main.go` builds your client.

If you change the configuration, run `make clean` before `make` to purge the generated files.

### Build with Docker

Use the `ghcr.io/fujiwara/aws-sdk-client-go:builder` builder image to build the client inside a container.

Environment variables:
- `GIT_REF`: Git reference to checkout the repository. Default is `main`. You can specify a branch, tag, or commit hash.

Example using the `AWS_SDK_CLIENT_GO_GEN` environment variable:
```console
$ docker run -it -e AWS_SDK_CLIENT_GO_GEN=ecs,firehose,s3 ghcr.io/fujiwara/aws-sdk-client-go:builder
...
```

Example using the `gen.yaml` configuration file:
```console
$ docker run -it -v $(pwd)/gen.yaml:/app/gen.yaml ghcr.io/fujiwara/aws-sdk-client-go:builder
...
Completed. Please extract /app/aws-sdk-client-go from this container!
For example, run the following command:
docker cp $(docker ps -lq):/app/aws-sdk-client-go .
```

After the build is complete, copy the binary to your host machine with the `docker cp` command.

```console
$ docker cp $(docker ps -lq):/app/aws-sdk-client-go .
```

### Docker Multi-stage build

It is also possible to use the `ghcr.io/fujiwara/aws-sdk-client-go:builder` builder image in a multi-stage build.

Run `./build-in-docker.sh` in the container to build the client. The built binary will be located in the `/app` directory. You can then copy it to the final image.

```Dockerfile
FROM ghcr.io/fujiwara/aws-sdk-client-go:builder AS builder
ENV AWS_SDK_CLIENT_GO_GEN=ecs,firehose,s3
ENV GIT_REF=v0.0.13
RUN ./build-in-docker.sh

FROM debian:bookworm-slim
COPY --from=builder /app/aws-sdk-client-go /usr/local/bin/aws-sdk-client-go
```

## Performance comparison

Example of executing `sts get-caller-identity` on a 0.25 vCPU Fargate(AMD64) using `/usr/bin/time -v` for time measurement.

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
  [<input>]      input JSON/Jsonnet struct or filename

Flags:
  -h, --help                      Show context-sensitive help.
  -i, --input-stream=STRING       bind input filename or '-' to io.Reader field in the input struct
  -o, --output-stream=STRING      bind output filename or '-' to io.ReadCloser field in the output struct
      --[no-]api-output           output API response into stdout
  -r, --raw-output                output raw strings, not JSON texts
  -c, --compact                   compact JSON output
  -q, --query=STRING              JMESPath query to apply to output
      --ext-str=KEY=VALUE;...     external variables for Jsonnet
      --ext-code=KEY=VALUE;...    external code for Jsonnet
      --[no-]strict               strict input JSON unmarshaling
  -f, --follow-next=""            OutputField=InputField format. follow the next token.
  -n, --dry-run                   dry-run mode
  -v, --version                   show version
```

- `service`: AWS service name.
- `method`: Method name of the service client.
- `input`: JSON input for the method.

The output is JSON format.

### Example usage

#### Show supported services

```console
$ aws-sdk-client-go
```

#### Show methods of the service

```console
$ aws-sdk-client-go ecs
```

#### Call method of the service

The third argument is a [JSON](https://json.org) or [Jsonnet](https://jsonnet.org/) input for the method. This can be omitted if the method requires no input (`{}` is passed implicitly).

```console
$ aws-sdk-client-go ecs DescribeClusters '{"Cluster":"default"}' # JSON
```

```console
$ aws-sdk-client-go ecs DescribeClusters "{Cluster:'default'}"   # Jsonnet
```

If the method name is "kebab-case", it automatically converts to "PascalCase" (i.e., `describe-clusters` -> `DescribeClusters`).

```console
$ aws-sdk-client-go ecs describe-clusters '{"Cluster":"default"}'
```

The third argument can be a filename that contains JSON or Jsonnet input.

```console
$ aws-sdk-client-go ecs DescribeClusters my.jsonnet
```

**Note**: By default, the input JSON is unmarshaled strictly. Unknown fields for the input struct in the input JSON cause an error. If you want to unmarshal the input JSON non-strictly, use `--no-strict` option.

#### `--ext-str` and `--ext-code` options

Pass external variables to Jsonnet.

This is useful when you want to use variables in Jsonnet.

```console
$ aws-sdk-client-go ecs DescribeClusters my.jsonnet --ext-str Cluster=default
```

```jsonnet
// my.jsonnet
{
  Cluster: std.extVar("Cluster"),
}
```

#### `--input-stream (-i)` option

Bind a file or stdin to the input struct.

```console
$ aws-sdk-client-go s3 put-object '{"Bucket": "my-bucket", "Key": "my.txt"}' --input-stream my.txt
```

[s3#PutObjectInput](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#PutObjectInput) has `Body` field of `io.Reader`. `--input-stream` option binds the file to the field.

When the input struct has only one field of `io.Reader`, `aws-sdk-client-go` reads the file and binds it to the field automatically. (Currently, all SDK input structs have at most one io.Reader field.)

When the input struct has a "\*Length" field for the size of the content, `aws-sdk-client-go` sets the size of the content to the field automatically. For example, [s3#PutObjectInput](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#PutObjectInput) has `ContentLength` field.

If `--input-stream` is "-", `aws-sdk-client-go` reads from stdin. In this case, `aws-sdk-client-go` reads all contents into memory, so it is not suitable for large files. Consider using a file for large content.

#### `--output-stream (-o)` option

Bind the `io.ReadCloser` of the API output to a file or stdout.

```console
$ aws-sdk-client-go s3 get-object '{"Bucket": "my-bucket", "Key": "my.txt"}' --output-stream my.txt
```

[s3#GetObjectOutput](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#PutObjectOutput) has `Body` field of `io.ReadeCloser`. `--output-stream` option binds the file to the field.

When the output struct has only one field of `io.ReadCloser`, `aws-sdk-client-go` copies it to the file automatically. (Currently, all SDK output structs have at most one io.ReadCloser field.)

If `--output-stream` is "-", `aws-sdk-client-go` writes into stdout. The result of the API also writes to stdout by default. If you don't want to output the result, use `--no-api-output`.

#### `--raw-output (-r)` option

Output raw strings, not JSON texts.

This option is like `jq -r`.

#### `--follow-next (-f)` option

Set the output/input field name of the next token. This option is useful for paginated APIs.

For example, [s3#ListObjectsV2Output](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#ListObjectsV2Output) has a `NextContinuationToken` field, and [s3#ListObjectsV2Input](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#ListObjectsV2Input) has a `ContinuationToken` field. You can follow the next token by the following command.

`{FieldInOutput}={FieldInInput}` format is used for `--follow-next` option.

```console
$ aws-sdk-client-go s3 list-objects-v2 '{"Bucket": "my-bucket"}' \
  --follow-next NextContinuationToken=ContinuationToken
```

If the same field name is used in the output and input, you can omit the input field name.

```console
$ aws-sdk-client-go ecs list-tasks '{"Cluster":"default"}' \
  --follow-next NextToken
```

#### Query output by JMESPath

Query the output by JMESPath like the AWS CLI.

```console
$ aws-sdk-client-go ecs DescribeClusters '{"Cluster":"default"}' \
  --query 'Clusters[0].ClusterArn'
```

#### Show help

For method-specific documentation, use the `help` argument to display the URL of the method's documentation. Since `aws-sdk-client-go` is a simple wrapper for the AWS SDK Go v2 service client, its usage mirros that of the SDK.

```console
$ aws-sdk-client-go ecs DescribeClusters help
See https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/ecs#Client.DescribeClusters
```

## LICENSE

MIT

## Author

Fujiwara Shunichiro

# aws-sdk-client-go

aws-sdk-client-go is a Go client for AWS SDK services.

This client is auto generated from the AWS SDK Go service client.

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

### Show methods of the service

```console
$ aws-sdk-client-go ecs
```

### Call method of the service

```console
$ aws-sdk-client-go ecs DescribeClusters '{"Cluster":"default"}'
```

The third argument is JSON input for the method. If the method does not require input, you can omit the third argument (implicitly `{}` passed).

```console

## LICENSE

MIT

## Author

Fujiwara Shunichiro

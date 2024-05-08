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

```console
$ aws-sdk-client-go ecs ListClusters
```

The third argument is optional. The value is passed to the method as a input parameter.

```console
$ aws-sdk-client-go ecs DescribeClusters '{"Cluster":"default"}'
```

## LICENSE

MIT

## Author

Fujiwara Shunichiro

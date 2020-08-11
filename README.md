![Terraform Module Tests](https://github.com/rodrigodelmonte/terraform-aws-kinesis-stream/workflows/Terraform%20Module%20Tests/badge.svg)

# AWS Kinesis Stream Terraform module

Terraform module which creates Kinesis Stream resources on AWS.

This type of resources are supported:

* [AWS Kinesis Stream](https://www.terraform.io/docs/providers/aws/r/kinesis_stream.html)

## Terraform version

Terraform 0.12. Module version to ~> v2.0.

## Usage

```hcl
module "kinesis-stream" {

  source  = "rodrigodelmonte/kinesis-stream/aws"
  version = "v2.0.1"

  name                      = "kinesis_stream_example"
  shard_count               = 1
  retention_period          = 24
  shard_level_metrics       = ["IncomingBytes", "OutgoingBytes"]
  enforce_consumer_deletion = false
  encryption_type           = "KMS"
  kms_key_id                = "alias/aws/kinesis"
  tags                      = {
      Name = "kinesis_stream_example"
  }

}
```

## Providers

| Name | Version |
|------|---------|
| aws  | ~> 2.23 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
|name|A name to identify the stream. This is unique to the AWS account and region the Stream is created in.|string||yes|
|shard_count|The number of shards that the stream will use|number|1|yes|
|retention_period|Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.|number|24|yes|
|shard_level_metrics|"A list of shard-level CloudWatch metrics which can be enabled for the stream.|list(string)|[]|no|
|enforce_consumer_deletion|A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error.|bool|false|no|
|encryption_type|The encryption type to use. The only acceptable values are NONE or KMS.|string|"NONE"|no|
|kms_key_id|The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias alias/aws/kinesis.|string|""|no|
|tags|A mapping of tags to assign to the resource|map|{}|no|

## Outputs

| Name | Description |
|------|-------------|
|kinesis_stream_name|The unique Kinesis stream name|
|kinesis_stream_shard_count|The count of shards for this Kinesis stream|
|kinesis_stream_arn|The Amazon Resource Name (ARN) specifying the Kinesis stream|
|kinesis_stream_iam_policy_read_only_arn|The IAM Policy (ARN) read only of the Stream|
|kinesis_stream_iam_policy_write_arn|The IAM Policy (ARN) write only of the Stream|
|kinesis_stream_iam_policy_admin_arn|The IAM Policy (ARN) admin of the Stream|

## Tests

This module has been tested using [Terratest](https://github.com/gruntwork-io/terratest)

```sh
# Install requirements
cd test/
go mod init
# Run tests
go test -v -timeout 30m
```

## Authors

Module is maintained by [Rodrigo Del Monte](https://github.com/rodrigodelmonte)

## License

Apache 2 Licensed. See LICENSE for full details.

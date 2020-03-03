resource "aws_kinesis_stream" "stream" {

  name                      = var.name
  shard_count               = var.shard_count
  retention_period          = var.retention_period
  shard_level_metrics       = var.shard_level_metrics
  enforce_consumer_deletion = var.enforce_consumer_deletion
  encryption_type           = var.encryption_type
  kms_key_id                = var.kms_key_id
  tags                      = var.tags

}

resource "aws_iam_policy" "read-only" {
  name        = format("kinesis-stream-%s-read-only", var.name)
  path        = "/"
  description = "Managed by Terraform"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = concat([
      {
        Effect = "Allow"
        Action = [
          "kinesis:DescribeLimits",
          "kinesis:DescribeStream",
          "kinesis:GetRecords",
          "kinesis:GetShardIterator",
          "kinesis:SubscribeToShard"
        ]
        Resource = [
          aws_kinesis_stream.stream.arn
        ]
      }
    ])
  })
}

resource "aws_iam_policy" "write-only" {
  name        = format("kinesis-stream-%s-write-only", var.name)
  path        = "/"
  description = "Managed by Terraform"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = concat([
      {
        Effect = "Allow"
        Action = [
          "kinesis:PutRecord",
          "kinesis:PutRecords",
        ]
        Resource = [
          aws_kinesis_stream.stream.arn
        ]
      }
    ])
  })
}

resource "aws_iam_policy" "admin" {
  name        = format("kinesis-stream-%s-admin", var.name)
  path        = "/"
  description = "Managed by Terraform"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = concat([
      {
        Effect = "Allow"
        Action = [
          "kinesis:*",
        ]
        Resource = [
          aws_kinesis_stream.stream.arn
        ]
      }
    ])
  })
}

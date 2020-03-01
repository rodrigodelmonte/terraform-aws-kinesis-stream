
provider "aws" {
  region  = "eu-west-1"
  version = "~> 2.48"
}

module "kinesis_stream_example" {
  source = "./../"

  name                      = var.name
  shard_count               = var.shard_count
  retention_period          = var.retention_period
  shard_level_metrics       = var.shard_level_metrics
  enforce_consumer_deletion = var.enforce_consumer_deletion
  encryption_type           = var.encryption_type
  kms_key_id                = var.kms_key_id
  tags                      = var.tags

}

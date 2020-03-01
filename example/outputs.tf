output "kinesis_stream_name" {
  description = "The unique Stream name "
  value       = module.kinesis_stream_example.kinesis_stream_name
}

output "kinesis_stream_shard_count" {
  description = "The count of Shards for this Stream"
  value       = module.kinesis_stream_example.kinesis_stream_shard_count
}

output "kinesis_stream_arn" {
  description = "The Amazon Resource Name (ARN) specifying the Stream"
  value       = module.kinesis_stream_example.kinesis_stream_arn
}

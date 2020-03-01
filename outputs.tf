output "kinesis_stream_name" {
  description = "The unique Stream name "
  value       = aws_kinesis_stream.stream.name
}

output "kinesis_stream_shard_count" {
  description = "The count of Shards for this Stream"
  value       = aws_kinesis_stream.stream.shard_count
}

output "kinesis_stream_arn" {
  description = "The Amazon Resource Name (ARN) specifying the Stream"
  value       = aws_kinesis_stream.stream.arn
}

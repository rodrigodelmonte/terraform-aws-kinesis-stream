package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestKinesisStream(t *testing.T) {
	t.Parallel()

	const expectedKinesisStreamName = "test"
	const expectedKinesisStreamShardCount = "1"
	const expectedKinesisStreamArn = "stream/test"
	const expectedKinesisStreamIamPolicyReadOnlyArn = "policy/kinesis-stream-test-read-only"
	const expectedKinesisStreamIamPolicyWriteOnlyArn = "policy/kinesis-stream-test-write-only"
	const expectedKinesisStreamIamPolicyAdminArn = "policy/kinesis-stream-test-admin"

	terraformOptions := &terraform.Options{
		// Source path of Terraform directory.
		TerraformDir: "../example/",
		Vars: map[string]interface{}{
			"name":                      expectedKinesisStreamName,
			"shard_count":               expectedKinesisStreamShardCount,
			"retention_period":          24,
			"shard_level_metrics":       []string{"IncomingBytes", "OutgoingBytes"},
			"enforce_consumer_deletion": false,
			"encryption_type":           "KMS",
			"kms_key_id":                "alias/aws/kinesis",
			"tags":                      map[string]string{"Name": "test"},
		},
	}

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// To clean up any resources that have been created, run 'terraform destroy' towards the end of the test
	defer terraform.Destroy(t, terraformOptions)

	kinesisStreamName := terraform.Output(t, terraformOptions, "kinesis_stream_name")
	kinesisStreamShardCount := terraform.Output(t, terraformOptions, "kinesis_stream_shard_count")
	kinesisStreamArn := terraform.Output(t, terraformOptions, "kinesis_stream_arn")
	kinesisStreamIamPolicyReadOnlyArn := terraform.Output(t, terraformOptions, "kinesis_stream_iam_policy_read_only_arn")
	kinesisStreamIamPolicyWriteOnlyArn := terraform.Output(t, terraformOptions, "kinesis_stream_iam_policy_write_only_arn")
	kinesisStreamIamPolicyAdminArn := terraform.Output(t, terraformOptions, "kinesis_stream_iam_policy_admin_arn")

	t.Run("Assert equal kinesis stream name", func(t *testing.T) {
		assert.Contains(t, kinesisStreamName, expectedKinesisStreamName)
	})

	t.Run("Assert equal kinesis shard count", func(t *testing.T) {
		assert.Contains(t, kinesisStreamShardCount, expectedKinesisStreamShardCount)
	})

	t.Run("Assert contains kinesis stream arn", func(t *testing.T) {
		assert.Contains(t, kinesisStreamArn, expectedKinesisStreamArn)
	})

	t.Run("Assert contains iam policy read only arn", func(t *testing.T) {
		assert.Contains(t, kinesisStreamIamPolicyReadOnlyArn, expectedKinesisStreamIamPolicyReadOnlyArn)
	})

	t.Run("Assert contains iam policy write only arn", func(t *testing.T) {
		assert.Contains(t, kinesisStreamIamPolicyWriteOnlyArn, expectedKinesisStreamIamPolicyWriteOnlyArn)
	})

	t.Run("Assert contains iam policy admin arn", func(t *testing.T) {
		assert.Contains(t, kinesisStreamIamPolicyAdminArn, expectedKinesisStreamIamPolicyAdminArn)
	})

}

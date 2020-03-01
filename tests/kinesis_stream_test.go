package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestKinesisStream(t *testing.T) {
	t.Parallel()

	const expectedKinesisStreamName = "kinesis_stream_test"
	const expectedKinesisStreamShardCount = "1"
	const expectedKinesisStreamArn = "stream/kinesis_stream_test"

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
			"tags":                      map[string]string{"Name": "kinesis_stream_test"},
		},
	}

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// To clean up any resources that have been created, run 'terraform destroy' towards the end of the test
	defer terraform.Destroy(t, terraformOptions)

	kinesisStreamName := terraform.Output(t, terraformOptions, "kinesis_stream_name")
	kinesisStreamShardCount := terraform.Output(t, terraformOptions, "kinesis_stream_shard_count")
	kinesisStreamArn := terraform.Output(t, terraformOptions, "kinesis_stream_arn")

	if expectedKinesisStreamName != kinesisStreamName {
		t.Errorf("Expected %s got %s", expectedKinesisStreamName, kinesisStreamName)
	}

	if expectedKinesisStreamShardCount != kinesisStreamShardCount {
		t.Errorf("Expected %s got %s", expectedKinesisStreamShardCount, kinesisStreamShardCount)
	}

	if strings.Contains(kinesisStreamArn, expectedKinesisStreamArn) != true {
		t.Errorf("Expected %s got %s", expectedKinesisStreamArn, kinesisStreamArn)
	}
}

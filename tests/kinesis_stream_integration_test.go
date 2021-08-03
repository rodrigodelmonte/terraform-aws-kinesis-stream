// +build integration_test

package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	kinesisStreamName       = "kinesis_stream_integration_test"
	kinesisStreamShardCount = "1"
)

var (
	streamName   = aws.String(kinesisStreamName)
	partitionKey = aws.String("key1")
	expectedData = []byte("test_abcd_123_the_test_string")
)

func TestKinesisStreamIntegrationTest(t *testing.T) {

	terraformOptions := &terraform.Options{
		// Source path of Terraform directory.
		TerraformDir: "../example/",
		Vars: map[string]interface{}{
			"name":                      kinesisStreamName,
			"shard_count":               kinesisStreamShardCount,
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

	// Load AWS credentials
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create Kinesis client
	client := kinesis.NewFromConfig(cfg)

	// shardId required to consume the record
	shardId := publishRecord(client, expectedData)

	// Wait for the stream acknowledge the record before consume record
	// TODO:
	// Evaluate usage of https://github.com/cenkalti/backoff
	time.Sleep(1 * time.Second)
	consumedData := consumeRecord(client, shardId)

	t.Run("Assert expected data is equal to consumed data", func(t *testing.T) {
		assert.Equal(t, string(expectedData), consumedData)
	})

}

func publishRecord(client *kinesis.Client, data []byte) string {

	putOutput, err := client.PutRecord(context.TODO(), &kinesis.PutRecordInput{
		Data:         data,
		StreamName:   streamName,
		PartitionKey: partitionKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	return *putOutput.ShardId
}

func consumeRecord(client *kinesis.Client, shardId string) string {

	var record string

	iteratorOutput, err := client.GetShardIterator(context.TODO(), &kinesis.GetShardIteratorInput{
		ShardId:           aws.String(shardId),
		ShardIteratorType: "TRIM_HORIZON",
		StreamName:        streamName,
	})
	if err != nil {
		log.Fatal(err)
	}

	shardIterator := iteratorOutput.ShardIterator

	for i := 0; i < 30; i++ {
		getOutput, err := client.GetRecords(context.TODO(), &kinesis.GetRecordsInput{
			ShardIterator: shardIterator,
			Limit:         aws.Int32(1),
		})
		if err != nil {
			log.Fatal(err)
		}

		if len(getOutput.Records) > 0 {
			record = string(getOutput.Records[0].Data)
			break
		}
		shardIterator = getOutput.NextShardIterator
		time.Sleep(1 * time.Second)
	}

	return record
}

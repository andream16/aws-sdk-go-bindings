package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestDynamoDBPutItemType struct {
	SomeParam string `json:"some_param"`
}

func TestDynamoDB_PutItem(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	svc, svcErr := New(awsSvc, cfg.DynamoDB.Endpoint)

	assert.NoError(t, svcErr)
	assert.NotEmpty(t, svc)

	testdata.MockDynamoDBTable(t, svc.DynamoDB.DynamoDB, cfg.DynamoDB.CmdTableName)

	in := new(TestDynamoDBPutItemType)
	in.SomeParam = "some_val"

	err := svc.PutItem(in, cfg.DynamoDB.CmdTableName)

	assert.NoError(t, err)

}

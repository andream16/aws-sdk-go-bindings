package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestDynamoDBPutItemType struct {
	SomeParam string `json:"some_param"`
}

func TestDynamoDB_Methods(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	testDynamoDBPutItem(t, cfg)
	testDynamoDBGetItem(t, cfg)

}

func testDynamoDBPutItem(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	svc, svcErr := New(awsSvc, cfg.DynamoDB.Endpoint)

	assert.NoError(t, svcErr)
	assert.NotEmpty(t, svc)

	testdata.MockDynamoDBTable(t, svc.DynamoDB.DynamoDB, cfg.DynamoDB.CmdTableName, cfg)

	in := new(TestDynamoDBPutItemType)
	in.SomeParam = "some_val"

	err := svc.PutItem(in, cfg.DynamoDB.CmdTableName)

	assert.NoError(t, err)

	shouldBeEmptyErr1 := svc.PutItem(in, "")

	assert.Error(t, shouldBeEmptyErr1)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr1.Error())

	shouldBeEmptyErr2 := svc.PutItem(struct{}{}, cfg.DynamoDB.CmdTableName)

	assert.Error(t, shouldBeEmptyErr2)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr2.Error())

}

func testDynamoDBGetItem(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	s := "some_val"

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	svc, svcErr := New(awsSvc, cfg.DynamoDB.Endpoint)

	assert.NoError(t, svcErr)
	assert.NotEmpty(t, svc)

	testdata.MockDynamoDBTable(t, svc.DynamoDB.DynamoDB, cfg.DynamoDB.CmdTableName, cfg)

	in := new(TestDynamoDBPutItemType)
	in.SomeParam = s

	err := svc.PutItem(in, cfg.DynamoDB.CmdTableName)

	assert.NoError(t, err)

	out, outErr := svc.GetItem(cfg.DynamoDB.CmdTableName, cfg.DynamoDB.PrimaryKey, s)

	assert.NoError(t, outErr)
	assert.NotEmpty(t, out)

	_, shouldBeEmptyErr1 := svc.GetItem("", cfg.DynamoDB.PrimaryKey, s)

	assert.Error(t, shouldBeEmptyErr1)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr1.Error())

	_, shouldBeEmptyErr2 := svc.GetItem(cfg.DynamoDB.CmdTableName, "", s)

	assert.Error(t, shouldBeEmptyErr2)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr2.Error())

	_, shouldBeEmptyErr3 := svc.GetItem(cfg.DynamoDB.CmdTableName, cfg.DynamoDB.PrimaryKey, "")

	assert.Error(t, shouldBeEmptyErr3)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr3.Error())

}

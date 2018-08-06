package dynamodb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

type TestDynamoDBDynamoPutItemType struct {
	SomeParam string `json:"some_param"`
}

func TestDynamoDB_Methods(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	testDynamoDBDynamoPutItem(t, cfg)
	testDynamoDBDynamoGetItem(t, cfg)

}

func testDynamoDBDynamoPutItem(t *testing.T, cfg configuration.Configuration) {

	dynamoSvc := testdata.MockDynamoDB(t, cfg)

	tableName := cfg.DynamoDB.PkgTableName

	testdata.MockDynamoDBTable(t, dynamoSvc, tableName, cfg)

	var input TestDynamoDBDynamoPutItemType
	input.SomeParam = cfg.DynamoDB.PrimaryKey

	putItemIn, putItemInErr := NewPutItemInput(input, tableName)

	assert.NoError(t, putItemInErr)

	dynamoNewSvc := new(DynamoDB)
	dynamoNewSvc.DynamoDB = dynamoSvc

	err := dynamoNewSvc.DynamoPutItem(putItemIn)

	assert.NoError(t, err)

}

func testDynamoDBDynamoGetItem(t *testing.T, cfg configuration.Configuration) {

	dynamoSvc := testdata.MockDynamoDB(t, cfg)

	tableName := cfg.DynamoDB.PkgTableName
	primaryKey := cfg.DynamoDB.PrimaryKey
	keyValue := cfg.DynamoDB.PrimaryKey

	testdata.MockDynamoDBTable(t, dynamoSvc, tableName, cfg)

	var input TestDynamoDBDynamoPutItemType
	input.SomeParam = cfg.DynamoDB.PrimaryKey

	putItemIn, putItemInErr := NewPutItemInput(input, tableName)

	assert.NoError(t, putItemInErr)

	dynamoNewSvc := new(DynamoDB)
	dynamoNewSvc.DynamoDB = dynamoSvc

	err := dynamoNewSvc.DynamoPutItem(putItemIn)

	assert.NoError(t, err)

	getItemInput, getItemInputErr := NewGetItemInput(tableName, primaryKey, keyValue)

	assert.NoError(t, getItemInputErr)

	getItemOut, err := dynamoNewSvc.DynamoGetItem(getItemInput)

	assert.NoError(t, err)
	assert.NotEmpty(t, getItemOut)

}

package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestDynamoDBDynamoPutItemType struct {
	SomeParam string `json:"some_param"`
}

func TestDynamoDB_DynamoPutItem(t *testing.T) {

	dynamoSvc := testdata.MockDynamoDB(t)

	cfg := testdata.MockConfiguration(t)

	testdata.MockDynamoDBTable(t, dynamoSvc, cfg.DynamoDB.PkgTableName)

	var input TestDynamoDBDynamoPutItemType
	input.SomeParam = cfg.DynamoDB.PrimaryKey

	putItemIn, putItemInErr := NewPutItemInput(input, cfg.DynamoDB.PkgTableName)

	assert.NoError(t, putItemInErr)

	dynamoNewSvc := new(DynamoDB)
	dynamoNewSvc.DynamoDB = dynamoSvc

	err := dynamoNewSvc.DynamoPutItem(putItemIn)

	assert.NoError(t, err)

}

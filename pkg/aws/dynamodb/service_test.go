package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestDynamoDBDynamoPutItemType struct {
	SomeParam string `json:"some_param"`
}

func TestDynamoDB_DynamoPutItem(t *testing.T) {

	dynamoSvc := testdata.MockDynamoDB(t)

	cfg := testdata.MockConfiguration(t)

	in := []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String(cfg.DynamoDB.PrimaryKey),
			AttributeType: aws.String("S"),
		},
	}

	testdata.CreateTableIfNotExists(
		t,
		*dynamoSvc,
		cfg.DynamoDB.TableName,
		in,
		[]*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(cfg.DynamoDB.PrimaryKey),
				KeyType:       aws.String("HASH"),
			},
		},
	)

	var input TestDynamoDBDynamoPutItemType
	input.SomeParam = cfg.DynamoDB.PrimaryKey

	putItemIn, putItemInErr := NewPutItemInput(input, cfg.DynamoDB.TableName)

	assert.NoError(t, putItemInErr)

	dynamoNewSvc := new(DynamoDB)
	dynamoNewSvc.DynamoDB = dynamoSvc

	err := dynamoNewSvc.DynamoPutItem(putItemIn)

	assert.NoError(t, err)

}

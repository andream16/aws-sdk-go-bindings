package testdata

import (
	"testing"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func CreateTableIfNotExists(t *testing.T, dynamo dynamodb.DynamoDB, tableName string, attributes []*dynamodb.AttributeDefinition, keys []*dynamodb.KeySchemaElement) {

	if !tableExists(t, dynamo, tableName) {
		createTable(t, dynamo, tableName, attributes, keys)
	}

}

func tableExists(t *testing.T, svc dynamodb.DynamoDB, tableName string) bool {

	t.Helper()

	_, err := svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return false
	}

	return true

}

func createTable(t *testing.T, dynamo dynamodb.DynamoDB, tableName string, attributes []*dynamodb.AttributeDefinition, keys []*dynamodb.KeySchemaElement) {

	t.Helper()

	_, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		TableName:            &tableName,
		AttributeDefinitions: attributes,
		KeySchema:            keys,
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	})

	assert.NoError(t, err)

}


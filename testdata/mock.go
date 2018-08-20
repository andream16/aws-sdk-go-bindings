package testdata

import (
	"testing"
	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	dynamodbStringType = "S"
	dynamodbHashType = "HASH"
)

func MockConfiguration(t *testing.T) *configuration.Configuration {

	t.Helper()

	cfg, err := configuration.Get()

	assert.NoError(t, err)
	assert.NotEmpty(t, cfg)

	return cfg

}

func MockDynamoDB(t *testing.T, cfg *configuration.Configuration) *dynamodb.DynamoDB {

	t.Helper()

	conf := &aws.Config{
		Region:   aws.String(cfg.Region),
		Endpoint: aws.String(cfg.DynamoDB.Endpoint),
	}

	dynamoSession, err := session.NewSession(conf)
	assert.NoError(t, err)

	return dynamodb.New(dynamoSession)

}

func MockDynamoDBTable(t *testing.T, svc *dynamodb.DynamoDB, tableName string, cfg *configuration.Configuration) {

	t.Helper()

	in := []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String(cfg.DynamoDB.PrimaryKey),
			AttributeType: aws.String(dynamodbStringType),
		},
	}

	CreateTableIfNotExists(
		t,
		*svc,
		tableName,
		in,
		[]*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(cfg.DynamoDB.PrimaryKey),
				KeyType:       aws.String(dynamodbHashType),
			},
		},
	)

}

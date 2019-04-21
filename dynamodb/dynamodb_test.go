package dynamodb

import (
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/pkg/errors"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (mockDynamoDBClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

type mockFailingDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (mockFailingDynamoDBClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, errors.New("some error")
}

func TestDynamoDB_PutItem(t *testing.T) {

	t.Run("should return an error because table name is empty", func(t *testing.T) {

		dynamoDB := &DynamoDB{}

		err := dynamoDB.PutItem("", nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	// This case is left untested since dynamodbattribute.MarshalMap does not return any error in aws-sdk-go v1.19.14
	t.Run("should return an error because there was an error marshaling item into map", func(t *testing.T) {})

	t.Run("should return an error because there was an error putting item", func(t *testing.T) {

		dynamoDB := &DynamoDB{
			dynamoDB: &mockFailingDynamoDBClient{},
		}

		err := dynamoDB.PutItem("some item", map[string]string{
			"some param": "some value",
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}

	})

	t.Run("should successfully put item", func(t *testing.T) {

		dynamoDB := &DynamoDB{
			dynamoDB: &mockDynamoDBClient{},
		}

		err := dynamoDB.PutItem("some item", map[string]string{
			"some param": "some value",
		})
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

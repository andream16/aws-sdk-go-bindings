package dynamodb

import (
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/pkg/errors"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (mockDynamoDBClient) PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func (mockDynamoDBClient) GetItem(in *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"keyName": &dynamodb.AttributeValue{
				S: aws.String("someValue"),
			},
		},
	}, nil
}

type mockFailingDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (mockFailingDynamoDBClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, errors.New("some error")
}

func (mockFailingDynamoDBClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, errors.New("some error")
}

type item struct {
	KeyName string
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

		err := dynamoDB.PutItem("someItem", map[string]string{
			"someKey": "someValue",
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}

	})

	t.Run("should successfully put item", func(t *testing.T) {

		dynamoDB := &DynamoDB{
			dynamoDB: &mockDynamoDBClient{},
		}

		err := dynamoDB.PutItem("someItem", map[string]string{
			"someKey": "someValue",
		})
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

func TestDynamoDB_GetItem(t *testing.T) {

	t.Run("should return an error because table name is empty", func(t *testing.T) {

		dynamoDB := &DynamoDB{}

		err := dynamoDB.GetItem("", "", "", nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because key name is empty", func(t *testing.T) {

		dynamoDB := &DynamoDB{}

		err := dynamoDB.GetItem("tableName", "", "", nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because value is empty", func(t *testing.T) {

		dynamoDB := &DynamoDB{}

		err := dynamoDB.GetItem("tableName", "keyName", "", nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because there was an error getting the item", func(t *testing.T) {

		dynamoDB := &DynamoDB{
			dynamoDB: &mockFailingDynamoDBClient{},
		}

		err := dynamoDB.GetItem("tableName", "keyName", "someValue", nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("should return an error because there was an error marshaling map into struct", func(t *testing.T) {

		dynamoDB := &DynamoDB{
			dynamoDB: &mockDynamoDBClient{},
		}

		it := item{}

		err := dynamoDB.GetItem("tableName", "keyName", "someValue", it)
		if err == nil {
			t.Fatal("expected error, got nil")
		}

	})

	t.Run("should successfully get item", func(t *testing.T) {

		dynamoDB := &DynamoDB{
			dynamoDB: &mockDynamoDBClient{},
		}

		it := item{}

		err := dynamoDB.GetItem("tableName", "keyName", "someValue", &it)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

		if it.KeyName == "" {
			t.Fatal("struct not filled")
		}

	})

}

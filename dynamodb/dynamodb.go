package dynamodb

import (
	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/pkg/errors"
)

// DynamoDBer describes DynamoDB API
type DynamoDBer interface {
	PutItem(table string, item interface{}) error
	GetItem(table, key, value string, out interface{}) ([]byte, error)
}

// DynamoDB is the alias for dynamodb
type DynamoDB struct {
	dynamoDB dynamodbiface.DynamoDBAPI
}

// New returns a new DynamoDB
func New(config *aws.Config) (*DynamoDB, error) {

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize dynamodb session")
	}

	return &DynamoDB{
		dynamoDB: dynamodb.New(sess),
	}, nil
}

// PutItem writes an item in the given DynamoDB table
func (db DynamoDB) PutItem(table string, item interface{}) error {

	if table == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "table")
	}

	itemMap, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return errors.Wrap(err, "error marshaling item into map")
	}

	if _, err := db.dynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      itemMap,
	}); err != nil {
		return errors.Wrap(err, "error putting item")
	}

	return nil
}

// GetItem reads from table the element having given primary key equal to given value
func (db DynamoDB) GetItem(table string, key string, value string, out interface{}) error {

	if table == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "table")
	}

	if key == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "key")
	}

	if value == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "value")
	}

	getOut, err := db.dynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
	})
	if err != nil {
		return errors.Wrapf(err, "error getting item with %s=%s", key, value)
	}

	if err := dynamodbattribute.UnmarshalMap(getOut.Item, out); err != nil {
		return errors.Wrap(err, "error unmarshaling map into struct")
	}

	return nil
}

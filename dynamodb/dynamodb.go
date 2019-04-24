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
	GetItem(table, key, value string, out interface{}) error
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

// PutItem add an item into a table
func (db DynamoDB) PutItem(table string, item interface{}) error {

	if table == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "table")
	}

	m, _ := dynamodbattribute.MarshalMap(item)

	if _, err := db.dynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      m,
	}); err != nil {
		return errors.Wrapf(err, "uneble to insert item in table %s", table)
	}

	return nil
}

// GetItem reads from table the element having given primary key equal to given value
// out has to be an pointer
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

	item, err := db.dynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
	})
	if err != nil {
		return errors.Wrapf(err, "unable to get item with %s=%s from table %s", key, value, table)
	}

	dynamodbattribute.UnmarshalMap(item.Item, out)

	return nil
}

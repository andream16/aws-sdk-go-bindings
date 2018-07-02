package dynamodb

import (
	"errors"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/dynamodb"
	"reflect"
)

// PutItem puts a passed input into a passed table
func (svc *DynamoDB) PutItem(input interface{}, tableName string) error {

	if reflect.DeepEqual(input, reflect.Zero(reflect.TypeOf(input)).Interface()) || tableName == "" {
		return errors.New(ErrEmptyParameter)
	}

	in, err := dynamodb.NewPutItemInput(input, tableName)
	if err != nil {
		return err
	}

	putErr := svc.DynamoPutItem(in)
	if putErr != nil {
		return putErr
	}

	return nil

}

// GetItem gets an item from dynamodb
func (svc *DynamoDB) GetItem(tableName, keyName, keyValue string) (*GetItemOutPut, error) {

	if tableName == "" || keyName == "" || keyValue == "" {
		return nil, errors.New(ErrEmptyParameter)
	}

	in := dynamodb.NewGetItemInput(
		tableName,
		keyName,
		keyValue,
	)

	getItemOut, getItemErr := svc.DynamoGetItem(in)
	if getItemErr != nil {
		return nil, getItemErr
	}

	out := new(GetItemOutPut)

	itemErr := dynamodb.UnmarshalGetItemOutput(getItemOut, &out)
	if itemErr != nil {
		return nil, itemErr
	}

	return out, nil

}

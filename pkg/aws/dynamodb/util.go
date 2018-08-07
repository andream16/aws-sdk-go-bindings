package dynamodb

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// NewPutItemInput returns a new *PutItemInput
func NewPutItemInput(in interface{}, tableName string) (*PutItemInput, error) {

	if len(tableName) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	dynamoInput, dynamoInputErr := dynamodbattribute.MarshalMap(in)
	if dynamoInputErr != nil {
		return nil, dynamoInputErr
	}

	putItemInput := new(dynamodb.PutItemInput)
	putItemInput = putItemInput.SetItem(dynamoInput)
	putItemInput = putItemInput.SetTableName(tableName)

	out := new(PutItemInput)
	out.PutItemInput = putItemInput

	return out, nil

}

// NewGetItemInput returns a new *GetItemInput
func NewGetItemInput(tableName, keyName, keyValue string) (*GetItemInput, error) {

	if len(tableName) == 0 || len(keyName) == 0 || len(keyValue) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	getItemInput := new(dynamodb.GetItemInput)
	getItemInput = getItemInput.SetTableName(tableName)
	getItemInput = getItemInput.SetKey(
		map[string]*dynamodb.AttributeValue{
			keyName: {
				S: aws.String(keyValue),
			},
		},
	)

	out := new(GetItemInput)
	out.GetItemInput = getItemInput

	return out, nil

}

// UnmarshalStreamImage unmarshals a dynamo stream image in a pointer to an interface
func UnmarshalStreamImage(in map[string]events.DynamoDBAttributeValue, out interface{}) error {

	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		return errors.New(ErrNoPointerParameter)
	}

	if len(in) == 0 {
		return errors.New(ErrEmptyMap)
	}

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range in {

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return marshalErr
		}

		var dbAttr dynamodb.AttributeValue

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr

	}

	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)

}

// UnmarshalGetItemOutput unmarshals a *GetItemOutput into a passed interface reference
func UnmarshalGetItemOutput(in *GetItemOutput, out interface{}) error {

	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		return errors.New(ErrNoPointerParameter)
	}

	unmarshalError := dynamodbattribute.UnmarshalMap(in.GetItemOutput.Item, out)
	if unmarshalError != nil {
		return unmarshalError
	}

	return nil

}

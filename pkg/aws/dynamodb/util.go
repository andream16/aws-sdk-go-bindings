package dynamodb

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// NewPutItemInput returns a new *PutItemInput
func NewPutItemInput(in interface{}, tableName string) (*PutItemInput, error) {

	dynamoInput, dynamoInputErr := dynamodbattribute.MarshalMap(in)
	if dynamoInputErr != nil {
		return nil, dynamoInputErr
	}

	out := &PutItemInput{
		&dynamodb.PutItemInput{
			Item:      dynamoInput,
			TableName: aws.String(tableName),
		},
	}

	return out, nil

}

// NewGetItemInput returns a new *GetItemInput
func NewGetItemInput(tableName, keyName, keyValue string) *GetItemInput {

	in := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyName: {
				S: aws.String(keyValue),
			},
		},
	}

	out := new(GetItemInput)
	out.GetItemInput = in

	return out

}

// UnmarshalStreamImage unmarshals a dynamo stream image in a pointer to an interface
func UnmarshalStreamImage(in map[string]events.DynamoDBAttributeValue, out interface{}) error {

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

	unmarshalError := dynamodbattribute.UnmarshalMap(in.GetItemOutput.Item, out)
	if unmarshalError != nil {
		return unmarshalError
	}

	return nil

}

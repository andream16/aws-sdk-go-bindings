package dynamodb

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

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

package dynamodb

import (
	"errors"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/dynamodb"
	"github.com/aws/aws-lambda-go/events"
	"reflect"
)

// GetItemOutPut contains information about an item retrieved from DynamoDB
type GetItemOutPut interface{}

// UnmarshalStreamImage unmarshals a new Image Coming from the event into a passed interface.
// Pass a reference as output.
func UnmarshalStreamImage(input events.DynamoDBEventRecord, output interface{}) error {

	if reflect.DeepEqual(input, reflect.Zero(reflect.TypeOf(input)).Interface()) {
		return errors.New(ErrEmptyParameter)
	}

	if reflect.ValueOf(output).Kind() != reflect.Ptr {
		return errors.New(ErrNoPointerParameter)
	}

	err := dynamodb.UnmarshalStreamImage(
		input.Change.NewImage,
		output,
	)
	if err != nil {
		return err
	}

	return nil

}

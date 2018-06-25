package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/dynamodb"
	"github.com/aws/aws-lambda-go/events"
)

// GetItemOutPut contains information about an item retrieved from DynamoDB
type GetItemOutPut interface{}

// UnmarshalStreamImage unmarshals a new Image Coming from the event into a passed interface.
// Pass a reference as output.
func UnmarshalStreamImage(input events.DynamoDBEventRecord, output interface{}) (err error) {

	err = dynamodb.UnmarshalStreamImage(
		input.Change.NewImage,
		output,
	)

	return

}

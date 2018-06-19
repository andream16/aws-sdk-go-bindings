package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/dynamodb"
	"github.com/aws/aws-lambda-go/events"
)

// UnmarshalStreamImage unmarshals a new Image Coming from the event into a passed interface.
// Pass a reference as output.
func UnmarshalStreamImage(input events.DynamoDBEventRecord, output interface{}) (err error) {

	err = dynamodb.UnmarshalStreamImage(
		input.Change.NewImage,
		output,
	)

	return

}

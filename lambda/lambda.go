package lambda

import (
	"encoding/json"
	"reflect"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

// UnmarshalDynamoEvent unmarshals a events.DynamoDBEventRecord into a given target.
// Out has to be a pointer.
func UnmarshalDynamoEvent(event events.DynamoDBEventRecord, out interface{}) error {

	if reflect.DeepEqual(event, reflect.Zero(reflect.TypeOf(event)).Interface()) {
		return errors.Wrap(bindings.ErrInvalidParameter, "event")
	}
	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		return errors.Wrap(bindings.ErrInvalidParameter, "out")
	}

	img := event.Change.NewImage
	if len(img) == 0 {
		return errors.New("event's image is empty")
	}

	m := make(map[string]*dynamodb.AttributeValue, len(img))

	for k, v := range img {

		b, err := v.MarshalJSON()
		if err != nil {
			return errors.Wrap(err, "unable to marshal current element to json")
		}

		var attr dynamodb.AttributeValue

		err = json.Unmarshal(b, &attr)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal current element to json")
		}

		m[k] = &attr

	}

	return dynamodbattribute.UnmarshalMap(m, out)

}

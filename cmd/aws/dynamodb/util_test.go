package dynamodb

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestUnmarshalStreamImageType struct {
	SomeParam string `json:"some_param"`
}

func TestUnmarshalStreamImage(t *testing.T) {

	var in events.DynamoDBAttributeValue

	someVal := "some_val"

	mock := []byte(`
		{ 
      		"M": {
				"some_param" : {
					"S" : "some_val"
				}
 			}
		}
	`)

	err := json.Unmarshal(mock, &in)
	assert.NoError(t, err)

	m := in.Map()

	event := events.DynamoDBEventRecord{
		Change: events.DynamoDBStreamRecord{
			NewImage: m,
		},
	}

	var out TestUnmarshalStreamImageType

	unmarshalErr := UnmarshalStreamImage(event, &out)

	assert.NoError(t, unmarshalErr)
	assert.Equal(t, someVal, out.SomeParam)

}

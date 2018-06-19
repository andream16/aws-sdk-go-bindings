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

	input := []byte(`
        { "M": 
            {
                "some_param": { "S": "Joe" }
            }
        }`)

	var av events.DynamoDBAttributeValue

	err := json.Unmarshal(input, &av)
	assert.NoError(t, err)

	mock := new(TestUnmarshalStreamImageType)

	m := av.Map()

	unmarshalErr := UnmarshalStreamImage(m, &mock)
	assert.NoError(t, unmarshalErr)

	assert.NotEmpty(t, mock)
	assert.Equal(t, "Joe", mock.SomeParam)

}

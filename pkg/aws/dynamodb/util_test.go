package dynamodb

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestUnmarshalStreamImageType struct {
	SomeParam string `json:"some_param"`
}

func TestNewPutItemInput(t *testing.T) {

	tableName := "some_name"
	in := new(TestUnmarshalStreamImageType)
	in.SomeParam = "some_val"

	out, err := NewPutItemInput(in, tableName)

	assert.NoError(t, err)
	assert.Equal(t, tableName, *out.TableName)

}

func TestNewGetItemInput(t *testing.T) {

	tableName := "some_table"
	keyName := "some_key"
	keyValue := "some_key_value"

	out := NewGetItemInput(
		tableName,
		keyName,
		keyValue,
	)

	assert.Equal(t, tableName, *out.TableName)
	assert.Equal(t, keyValue, *out.Key[keyName].S)

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

func TestUnmarshalGetItemOutput(t *testing.T) {

	s := "Joe"

	var out TestUnmarshalStreamImageType

	getItemIn := &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"some_param": {
				S: aws.String(s),
			},
		},
	}

	in := new(GetItemOutput)
	in.GetItemOutput = getItemIn

	err := UnmarshalGetItemOutput(in, &out)

	assert.NoError(t, err)
	assert.Equal(t, s, out.SomeParam)

}

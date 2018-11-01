package dynamodb

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

type TestUnmarshalStreamImageType struct {
	SomeParam string `json:"some_param"`
}

func TestNewPutItemInput(t *testing.T) {

	tableName := "some_name"
	in := &TestUnmarshalStreamImageType{
		SomeParam: "some_val",
	}

	out, err := NewPutItemInput(in, tableName)

	assert.NoError(t, err)
	assert.Equal(t, tableName, *out.TableName)

	_, err = NewPutItemInput(in, "")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

}

func TestNewGetItemInput(t *testing.T) {

	tableName := "some_table"
	keyName := "some_key"
	keyValue := "some_key_value"

	out, err := NewGetItemInput(
		tableName,
		keyName,
		keyValue,
	)

	assert.NoError(t, err)
	assert.Equal(t, tableName, *out.TableName)
	assert.Equal(t, keyValue, *out.Key[keyName].S)

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

	err = UnmarshalStreamImage(event, &out)
	assert.NoError(t, err)

	assert.NotEmpty(t, out)
	assert.Equal(t, someVal, out.SomeParam)

	err = UnmarshalStreamImage(event, TestUnmarshalStreamImageType{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrNoPointerParameter)

	err = UnmarshalStreamImage(events.DynamoDBEventRecord{}, &mock)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

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

	err := UnmarshalGetItemOutput(getItemIn, &out)

	assert.NoError(t, err)
	assert.Equal(t, s, out.SomeParam)

	err = UnmarshalGetItemOutput(getItemIn, TestUnmarshalStreamImageType{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrNoPointerParameter)

}

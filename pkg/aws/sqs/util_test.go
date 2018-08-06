package sqs

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestNewCreateQueueInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)
	
	out, err := NewCreateQueueInput(cfg.SQS.QueueName)

	assert.NoError(t, err)
	assert.Equal(t, cfg.SQS.QueueName, *out.CreateQueueInput.QueueName)

	_, shouldBeErrEmptyParameter := NewCreateQueueInput("")

	assert.Error(t, shouldBeErrEmptyParameter)
	assert.Equal(t, ErrEmptyParameter, shouldBeErrEmptyParameter.Error())

}

func TestNewGetQueueAttributesInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetQueueAttributesInput(cfg.SQS.QueueName)

	assert.NoError(t, err)
	assert.Equal(t, cfg.SQS.QueueName, *out.GetQueueAttributesInput.QueueUrl)

	_, shouldBeEmptyErr := NewGetQueueAttributesInput("")

	assert.Error(t, shouldBeEmptyErr)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr.Error())

}

func TestNewSendMessageInput(t *testing.T) {

	queueUrl := "some_url"
	s := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	out, err := NewSendMessageInput(
		s,
		queueUrl,
	)

	assert.NoError(t, err)
	assert.Equal(t, queueUrl, *out.SendMessageInput.QueueUrl)

	b, bErr := base64.StdEncoding.DecodeString(*out.SendMessageInput.MessageBody)

	assert.NoError(t, bErr)
	assert.NotEmpty(t, b, val1, val2)

	var m TestSQSUtilType

	unmarshalErr := json.Unmarshal([]byte(b), &m)

	assert.NoError(t, unmarshalErr)
	assert.Equal(t, val1, m.SomeParam1)
	assert.Equal(t, val2, m.SomeParam2)

	_, shouldBeErrEmptyParameter := NewSendMessageInput(
		s,
		"",
	)

	assert.Error(t, shouldBeErrEmptyParameter)
	assert.Equal(t, ErrEmptyParameter, shouldBeErrEmptyParameter.Error())

	_, shouldBeErrNoPointerParameterAllowed := NewSendMessageInput(
		&s,
		queueUrl,
	)

	assert.Error(t, shouldBeErrNoPointerParameterAllowed)
	assert.Equal(t, ErrNoPointerParameterAllowed, shouldBeErrNoPointerParameterAllowed.Error())

}

func TestNewGetQueueUrlInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetQueueUrlInput(cfg.SQS.QueueName)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)
	assert.NotEmpty(t, out.QueueName)

}

func TestMarshalStructToJsonString(t *testing.T) {

	s := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	res, err := marshalStructToJson(s)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Contains(t, string(res), val1, val2)

	var out TestSQSUtilType

	unmarshalErr := json.Unmarshal([]byte(res), &out)

	assert.NoError(t, unmarshalErr)
	assert.Equal(t, val1, out.SomeParam1)
	assert.Equal(t, val2, out.SomeParam2)

	_, shouldBeErrNoPointerParameterAllowed := marshalStructToJson(&s)

	assert.Error(t, shouldBeErrNoPointerParameterAllowed)
	assert.Equal(t, ErrNoPointerParameterAllowed, shouldBeErrNoPointerParameterAllowed.Error())

}

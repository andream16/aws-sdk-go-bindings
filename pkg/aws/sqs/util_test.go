package sqs

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
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

	base64Out, base64Err := NewSendMessageInput(
		s,
		queueUrl,
		true,
	)

	assert.NoError(t, base64Err)
	assert.Equal(t, queueUrl, *base64Out.SendMessageInput.QueueUrl)

	b, bErr := base64.StdEncoding.DecodeString(*base64Out.SendMessageInput.MessageBody)

	assert.NoError(t, bErr)
	assert.NotEmpty(t, b)

	var m1 TestSQSUtilType

	unmarshalBase64OutErr := json.Unmarshal([]byte(b), &m1)

	assert.NoError(t, unmarshalBase64OutErr)
	assert.Equal(t, val1, m1.SomeParam1)
	assert.Equal(t, val2, m1.SomeParam2)

	out, err := NewSendMessageInput(
		s,
		queueUrl,
		false,
	)

	assert.NoError(t, err)
	assert.Equal(t, queueUrl, *out.SendMessageInput.QueueUrl)

	var m2 TestSQSUtilType

	unmarshalOutErr := json.Unmarshal([]byte(b), &m2)

	assert.NoError(t, unmarshalOutErr)
	assert.Equal(t, val1, m2.SomeParam1)
	assert.Equal(t, val2, m2.SomeParam2)

	_, shouldBeErrEmptyParameter := NewSendMessageInput(
		s,
		"",
		true,
	)

	assert.Error(t, shouldBeErrEmptyParameter)
	assert.Equal(t, ErrEmptyParameter, shouldBeErrEmptyParameter.Error())

	_, shouldBeErrNoPointerParameterAllowed := NewSendMessageInput(
		&s,
		queueUrl,
		true,
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

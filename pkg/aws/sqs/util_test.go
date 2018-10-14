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
	assert.Equal(t, cfg.SQS.QueueName, *out.QueueName)

	_, shouldBeErrEmptyParameter := NewCreateQueueInput("")

	assert.Error(t, shouldBeErrEmptyParameter)
	assert.Contains(t, shouldBeErrEmptyParameter.Error(), ErrEmptyParameter)

}

func TestNewGetQueueAttributesInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetQueueAttributesInput(cfg.SQS.QueueUrl)

	assert.NoError(t, err)
	assert.Equal(t, cfg.SQS.QueueUrl, *out.QueueUrl)

	_, shouldBeEmptyErr := NewGetQueueAttributesInput("")

	assert.Error(t, shouldBeEmptyErr)
	assert.Contains(t, shouldBeEmptyErr.Error(), ErrEmptyParameter)

}

func TestNewSendMessageInput(t *testing.T) {

	queueUrl := "some_url"
	s := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	out, outErr := NewSendMessageInput(
		s,
		queueUrl,
		true,
	)

	assert.NoError(t, outErr)
	assert.Equal(t, queueUrl, *out.QueueUrl)

	stringBody, stringBodyErr := base64.StdEncoding.DecodeString(*out.MessageBody)

	assert.NoError(t, stringBodyErr)
	assert.NotEmpty(t, stringBody)

	var m1 TestSQSUtilType

	unmarshalBase64OutErr := json.Unmarshal([]byte(stringBody), &m1)

	assert.NoError(t, unmarshalBase64OutErr)
	assert.Equal(t, val1, m1.SomeParam1)
	assert.Equal(t, val2, m1.SomeParam2)

	out, err := NewSendMessageInput(
		s,
		queueUrl,
		false,
	)

	assert.NoError(t, err)
	assert.Equal(t, queueUrl, *out.QueueUrl)

	var m2 TestSQSUtilType

	unmarshalOutErr := json.Unmarshal([]byte(stringBody), &m2)

	assert.NoError(t, unmarshalOutErr)
	assert.Equal(t, val1, m2.SomeParam1)
	assert.Equal(t, val2, m2.SomeParam2)

	_, shouldBeErrEmptyParameter := NewSendMessageInput(
		s,
		"",
		true,
	)

	assert.Error(t, shouldBeErrEmptyParameter)
	assert.Contains(t, shouldBeErrEmptyParameter.Error(), ErrEmptyParameter)

	_, shouldBeErrNoPointerParameterAllowed := NewSendMessageInput(
		&s,
		queueUrl,
		true,
	)

	assert.Error(t, shouldBeErrNoPointerParameterAllowed)
	assert.Contains(t, shouldBeErrNoPointerParameterAllowed.Error(), ErrNoPointerParameterAllowed)

}

func TestNewGetQueueUrlInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetQueueUrlInput(cfg.SQS.QueueName)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)
	assert.NotEmpty(t, out.QueueName)

}

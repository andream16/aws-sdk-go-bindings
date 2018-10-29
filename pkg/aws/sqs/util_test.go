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

	_, err = NewCreateQueueInput("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

}

func TestNewGetQueueAttributesInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetQueueAttributesInput(cfg.SQS.QueueUrl)

	assert.NoError(t, err)
	assert.Equal(t, cfg.SQS.QueueUrl, *out.QueueUrl)

	_, err = NewGetQueueAttributesInput("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

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
		true,
	)

	assert.NoError(t, err)
	assert.Equal(t, queueUrl, *out.QueueUrl)

	stringBody, err := base64.StdEncoding.DecodeString(*out.MessageBody)

	assert.NoError(t, err)
	assert.NotEmpty(t, stringBody)

	var m1 TestSQSUtilType

	err = json.Unmarshal([]byte(stringBody), &m1)

	assert.NoError(t, err)
	assert.Equal(t, val1, m1.SomeParam1)
	assert.Equal(t, val2, m1.SomeParam2)

	out, err = NewSendMessageInput(
		s,
		queueUrl,
		false,
	)

	assert.NoError(t, err)
	assert.Equal(t, queueUrl, *out.QueueUrl)

	var m2 TestSQSUtilType

	err = json.Unmarshal([]byte(stringBody), &m2)

	assert.NoError(t, err)
	assert.Equal(t, val1, m2.SomeParam1)
	assert.Equal(t, val2, m2.SomeParam2)

	_, err = NewSendMessageInput(
		s,
		"",
		true,
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

	_, err = NewSendMessageInput(
		&s,
		queueUrl,
		true,
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrNoPointerParameterAllowed)

}

func TestNewGetQueueUrlInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetQueueUrlInput(cfg.SQS.QueueName)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)
	assert.NotEmpty(t, out.QueueName)

}

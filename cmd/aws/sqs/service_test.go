package sqs

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	queueName = "some_queue_name"
	val1      = "some_val1"
	val2      = "some_val2"
)

type TestSQSUtilType struct {
	SomeParam1 string `json:"some_param_1"`
	SomeParam2 string `json:"some_param_2"`
}

func TestSQS_CreateQueue(t *testing.T) {

	svc := mockSQSSvc(t)

	err := svc.CreateQueue(queueName)

	assert.NoError(t, err)

	shouldBeEmptyParameterErr := svc.CreateQueue("")

	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyParameterErr.Error())

}

func TestSQS_SendMessage(t *testing.T) {

	queueUrl := testdata.MockConfiguration(t).SQS.QueueUrl

	svc := mockSQSSvc(t)

	createQueueErr := svc.CreateQueue(queueName)

	assert.NoError(t, createQueueErr)

	in := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	err := svc.SendMessage(
		in,
		queueUrl,
	)

	assert.NoError(t, err)

	shouldBeEmptyParameterErr := svc.SendMessage(
		in,
		"",
	)

	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyParameterErr.Error())

	shouldBeNoPointerParameterErr := svc.SendMessage(
		&in,
		queueUrl,
	)

	assert.Equal(t, ErrNoPointerParameterAllowed, shouldBeNoPointerParameterErr.Error())

}

func TestSQS_GetQueueAttributes(t *testing.T) {

	queueUrl := testdata.MockConfiguration(t).SQS.QueueUrl

	svc := mockSQSSvc(t)

	createQueueErr := svc.CreateQueue(queueName)

	assert.NoError(t, createQueueErr)

	_, err := svc.GetQueueAttributes(queueUrl)

	assert.NoError(t, err)

	_, shouldBeEmptyParameterErr := svc.GetQueueAttributes("")

	assert.Error(t, shouldBeEmptyParameterErr)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyParameterErr.Error())

}

func mockSQSSvc(t *testing.T) *SQS {

	t.Helper()

	cfg := testdata.MockConfiguration(t)

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	svc, svcErr := New(awsSvc)

	assert.NoError(t, svcErr)
	assert.NotEmpty(t, svc)

	return svc

}

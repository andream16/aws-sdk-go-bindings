package sqs

import (
	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
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

func TestSQS_SQSCreateQueue(t *testing.T) {

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, queueName)

}

func TestSQS_SQSGetQueueAttributes(t *testing.T) {

	queueUrl := testdata.MockConfiguration(t).SQS.QueueUrl

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, queueName)

	getQueueAttrsIn, getQueueAttrsErr := NewGetQueueAttributesInput(queueUrl)

	assert.NoError(t, getQueueAttrsErr)
	assert.Equal(t, queueUrl, *getQueueAttrsIn.GetQueueAttributesInput.QueueUrl)

	err := svc.SQSGetQueueAttributes(
		getQueueAttrsIn,
	)

	assert.NoError(t, err)

	badQueueUrl := `https://sqs.eu-central-1.amazonaws.com/150285746666/some_queue_name`

	badGetQueueAttrsIn, _ := NewGetQueueAttributesInput(badQueueUrl)

	shouldBeErr := svc.SQSGetQueueAttributes(
		badGetQueueAttrsIn,
	)

	assert.Error(t, shouldBeErr)

}

func TestSQS_SQSSendMessage(t *testing.T) {

	queueUrl := testdata.MockConfiguration(t).SQS.QueueUrl

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, queueName)

	m := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	sendMsgIn, sendMsgInErr := NewSendMessageInput(
		m,
		queueUrl,
	)

	assert.NoError(t, sendMsgInErr)
	assert.Contains(t, *sendMsgIn.MessageBody, val1, val2)

	err := svc.SQSSendMessage(sendMsgIn)

	assert.NoError(t, err)

}

func createSQSQueue(t *testing.T, svc *SQS, queueName string) {

	t.Helper()

	createQueueIn, createQueueInErr := NewCreateQueueInput(queueName)

	assert.NoError(t, createQueueInErr)

	assert.Equal(t, queueName, *createQueueIn.CreateQueueInput.QueueName)

	err := svc.SQSCreateQueue(createQueueIn)

	assert.NoError(t, err)

}

func newSQSSvc(t *testing.T) *SQS {

	t.Helper()

	cfg := testdata.MockConfiguration(t)

	svcIn := pkgAws.NewSessionInput(cfg.Region)

	awsSvc, awsSvcErr := pkgAws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	sqsSvc, sqsSvcErr := New(awsSvc)

	assert.NoError(t, sqsSvcErr)
	assert.NotEmpty(t, sqsSvc)

	return sqsSvc

}

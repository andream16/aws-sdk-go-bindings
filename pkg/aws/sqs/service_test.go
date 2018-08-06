package sqs

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

const (
	val1      = "some_val1"
	val2      = "some_val2"
)

type TestSQSUtilType struct {
	SomeParam1 string `json:"some_param_1"`
	SomeParam2 string `json:"some_param_2"`
}

func TestSQS_SQSCreateQueue(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, cfg.SQS.QueueName)

}

func TestSQS_SQSGetQueueAttributes(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, cfg.SQS.QueueName)

	getQueueUrlIn, err := NewGetQueueUrlInput(cfg.SQS.QueueName)

	assert.NoError(t, err)

	url, urlErr := svc.SQSGetQueueUrl(getQueueUrlIn)

	assert.NoError(t, urlErr)
	assert.NotEmpty(t, url)

	getQueueAttrsIn, getQueueAttrsInErr := NewGetQueueAttributesInput(url)

	assert.NoError(t, getQueueAttrsInErr)
	assert.Equal(t, cfg.SQS.QueueUrl, *getQueueAttrsIn.GetQueueAttributesInput.QueueUrl)

	_, getQueueAttrsErr := svc.SQSGetQueueAttributes(
		getQueueAttrsIn,
	)

	assert.NoError(t, getQueueAttrsErr)

	badQueueUrl := `badURL`

	badGetQueueAttrsIn, _ := NewGetQueueAttributesInput(badQueueUrl)

	_, shouldBeErr := svc.SQSGetQueueAttributes(
		badGetQueueAttrsIn,
	)

	assert.Error(t, shouldBeErr)

}

func TestSQS_SQSSendMessage(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, cfg.SQS.QueueName)

	m := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	sendMsgIn, sendMsgInErr := NewSendMessageInput(
		m,
		cfg.SQS.QueueName,
	)

	assert.NoError(t, sendMsgInErr)

	b, bErr := base64.StdEncoding.DecodeString(*sendMsgIn.MessageBody)

	assert.NoError(t, bErr)

	var o TestSQSUtilType

	marshalErr := json.Unmarshal([]byte(b), &o)

	assert.NoError(t, marshalErr)
	assert.Equal(t, o.SomeParam1, val1)
	assert.Equal(t, o.SomeParam2, val2)

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

	svcIn, svcInErr := pkgAws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)

	awsSvc, awsSvcErr := pkgAws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	sqsSvc, sqsSvcErr := New(awsSvc, cfg.SQS.Endpoint)

	assert.NoError(t, sqsSvcErr)
	assert.NotEmpty(t, sqsSvc)

	return sqsSvc

}

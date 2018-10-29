package sqs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

const (
	val1 = "some_val1"
	val2 = "some_val2"
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

	url, err := svc.SQSGetQueueUrl(cfg.SQS.QueueName)

	assert.NoError(t, err)
	assert.NotEmpty(t, url)

	_, err = svc.SQSGetQueueAttributes(url)

	assert.NoError(t, err)

	badQueueUrl := `badURL`

	_, err = svc.SQSGetQueueAttributes(badQueueUrl)

	assert.Error(t, err)

}

func TestSQS_SQSSendMessage(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, cfg.SQS.QueueName)

	m := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	err := svc.SQSSendMessage(
		m,
		cfg.SQS.QueueName,
		true,
	)

	assert.NoError(t, err)

}

func TestSQS_SQSGetQueueUrl(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := newSQSSvc(t)

	createSQSQueue(t, svc, cfg.SQS.QueueName)

	url, err := svc.SQSGetQueueUrl(cfg.SQS.QueueName)

	assert.NoError(t, err)
	assert.NotEmpty(t, url)
	assert.Equal(t, cfg.SQS.QueueUrl, url)

	_, err = svc.SQSGetQueueUrl("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

}

func createSQSQueue(t *testing.T, svc *SQS, queueName string) {

	t.Helper()

	err := svc.SQSCreateQueue(queueName)

	assert.NoError(t, err)

	err = svc.SQSCreateQueue("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

}

func newSQSSvc(t *testing.T) *SQS {

	t.Helper()

	cfg := testdata.MockConfiguration(t)

	svcIn, err := pkgAws.NewSessionInput(cfg.Region)

	assert.NoError(t, err)

	awsSvc, err := pkgAws.New(svcIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, awsSvc)

	sqsSvc, err := New(awsSvc, cfg.SQS.Endpoint)

	assert.NoError(t, err)
	assert.NotEmpty(t, sqsSvc)

	return sqsSvc

}

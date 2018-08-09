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

	url, urlErr := svc.SQSGetQueueUrl(cfg.SQS.QueueName)

	assert.NoError(t, urlErr)
	assert.NotEmpty(t, url)

	_, getQueueAttrsErr := svc.SQSGetQueueAttributes(url)

	assert.NoError(t, getQueueAttrsErr)

	badQueueUrl := `badURL`

	_, shouldBeErr := svc.SQSGetQueueAttributes(badQueueUrl)

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

	url, urlErr := svc.SQSGetQueueUrl(cfg.SQS.QueueName)

	assert.NoError(t, urlErr)
	assert.NotEmpty(t, url)
	assert.Equal(t, cfg.SQS.QueueUrl, url)

	_, shouldBeEmptyErr := svc.SQSGetQueueUrl("")

	assert.Error(t, shouldBeEmptyErr)
	assert.Contains(t, shouldBeEmptyErr.Error(), ErrEmptyParameter)

}

func createSQSQueue(t *testing.T, svc *SQS, queueName string) {

	t.Helper()

	err := svc.SQSCreateQueue(queueName)

	assert.NoError(t, err)

	shouldBeEmptyParamErr := svc.SQSCreateQueue("")

	assert.Error(t, shouldBeEmptyParamErr)
	assert.Contains(t, shouldBeEmptyParamErr.Error(), ErrEmptyParameter)

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

package sqs

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/aws/aws-sdk-go/service/sqs"

	intError "github.com/andream16/aws-sdk-go-bindings/internal/error"
)

const (
	QueueName = "queueName"
	QueueUrl  = "queueUrl"
	Input     = "input"
)

// NewCreateQueueInput creates a new queue given its name
func NewCreateQueueInput(queueName string) (*sqs.CreateQueueInput, error) {

	if len(queueName) == 0 {
		return nil, intError.FormatError(QueueName, ErrEmptyParameter)
	}

	createQueueIn := new(sqs.CreateQueueInput)
	createQueueIn = createQueueIn.SetQueueName(queueName)

	out := new(sqs.CreateQueueInput)
	out = createQueueIn

	return out, nil

}

// NewGetQueueAttributesInput returns a new *sqs.GetQueueAttributesInput given a queueUrl
func NewGetQueueAttributesInput(queueUrl string) (*sqs.GetQueueAttributesInput, error) {

	if len(queueUrl) == 0 {
		return nil, intError.FormatError(QueueUrl, ErrEmptyParameter)
	}

	out := new(sqs.GetQueueAttributesInput)
	out = out.SetQueueUrl(queueUrl)

	return out, nil

}

// NewSendMessageInput returns a new *sqs.SendMessageInput initialized with queueUrl and messageBody.
// If base64Encode = true then the messageBody will be encoded in base64
func NewSendMessageInput(input interface{}, queueUrl string, base64Encode bool) (*sqs.SendMessageInput, error) {

	if reflect.DeepEqual(reflect.TypeOf(input).Kind(), reflect.Ptr) {
		return nil, intError.FormatError(Input, ErrNoPointerParameterAllowed)
	}

	if len(queueUrl) == 0 {
		return nil, intError.FormatError(QueueUrl, ErrEmptyParameter)
	}

	out := new(sqs.SendMessageInput)

	msgBody, err := marshalStructToJson(input)
	if err != nil {
		return nil, err
	}

	msgStringBody := string(msgBody)

	if base64Encode {
		msgStringBody = base64.StdEncoding.EncodeToString(msgBody)
	}

	out = out.SetMessageBody(msgStringBody)
	out = out.SetQueueUrl(queueUrl)

	return out, nil

}

// NewGetQueueUrlInput returns a new *sqs.GetQueueUrlInput given a queue name
func NewGetQueueUrlInput(queueName string) (*sqs.GetQueueUrlInput, error) {

	if len(queueName) == 0 {
		return nil, intError.FormatError(QueueName, ErrEmptyParameter)
	}

	out := new(sqs.GetQueueUrlInput)
	out = out.SetQueueName(queueName)

	return out, nil

}

// marshalStructToJson marshals input into a []byte contains its json encoding
func marshalStructToJson(input interface{}) ([]byte, error) {

	if reflect.DeepEqual(reflect.TypeOf(input).Kind(), reflect.Ptr) {
		return nil, errors.New(ErrNoPointerParameterAllowed)
	}

	b, marshalErr := json.Marshal(input)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return b, nil

}

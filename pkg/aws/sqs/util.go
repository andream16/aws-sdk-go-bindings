package sqs

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/aws/aws-sdk-go/service/sqs"

	intError "github.com/andream16/aws-sdk-go-bindings/internal/error"
)

// NewCreateQueueInput creates a new queue given its name
func NewCreateQueueInput(queueName string) (*sqs.CreateQueueInput, error) {

	if len(queueName) == 0 {
		return nil, intError.Format(QueueName, ErrEmptyParameter)
	}

	out := &sqs.CreateQueueInput{
		QueueName: queueName
	}

	return out, nil

}

// NewGetQueueAttributesInput returns a new *sqs.GetQueueAttributesInput given a queueUrl
func NewGetQueueAttributesInput(queueUrl string) (*sqs.GetQueueAttributesInput, error) {

	if len(queueUrl) == 0 {
		return nil, intError.Format(QueueUrl, ErrEmptyParameter)
	}

	out := &sqs.GetQueueAttributesInput{
		QueueUrl: queueUrl
	}

	return out, nil

}

// NewSendMessageInput returns a new *sqs.SendMessageInput initialized with queueUrl and messageBody.
// If base64Encode = true then the messageBody will be encoded in base64
func NewSendMessageInput(input interface{}, queueUrl string, base64Encode bool) (*sqs.SendMessageInput, error) {

	if reflect.DeepEqual(reflect.TypeOf(input).Kind(), reflect.Ptr) {
		return nil, intError.Format(Input, ErrNoPointerParameterAllowed)
	}

	if len(queueUrl) == 0 {
		return nil, intError.Format(QueueUrl, ErrEmptyParameter)
	}

	msgBody, err := marshalStructToJson(input)
	if err != nil {
		return nil, err
	}

	msgStringBody := string(msgBody)

	if base64Encode {
		msgStringBody = base64.StdEncoding.EncodeToString(msgBody)
	}

	out := &sqs.SendMessageInput{
		MessageBody: msgStringBody,
		QueueUrl: queueUrl
	}

	return out, nil

}

// NewGetQueueUrlInput returns a new *sqs.GetQueueUrlInput given a queue name
func NewGetQueueUrlInput(queueName string) (*sqs.GetQueueUrlInput, error) {

	if len(queueName) == 0 {
		return nil, intError.Format(QueueName, ErrEmptyParameter)
	}

	out := &sqs.GetQueueUrlInput{
		QueueName: queueName
	}

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

package sqs

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/sqs"
	"gitlab.easy-network.it/sender-alert/leprotto/pkg/json"
	"reflect"
)

// NewCreateQueueInput creates a new queue given its name
func NewCreateQueueInput(queueName string) (*CreateQueueInput, error) {

	if len(queueName) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	createQueueIn := new(sqs.CreateQueueInput)

	createQueueIn = createQueueIn.SetQueueName(queueName)

	out := new(CreateQueueInput)
	out.CreateQueueInput = createQueueIn

	return out, nil

}

// NewGetQueueAttributesInput returns a new *GetQueueAttributesInput given a queueUrl
func NewGetQueueAttributesInput(queueUrl string) (*GetQueueAttributesInput, error) {

	if len(queueUrl) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	getQueueAttrsIn := new(sqs.GetQueueAttributesInput)
	getQueueAttrsIn = getQueueAttrsIn.SetQueueUrl(queueUrl)

	out := new(GetQueueAttributesInput)
	out.GetQueueAttributesInput = getQueueAttrsIn

	return out, nil

}

// NewSendMessageInput returns a new *SendMessageInput initialized with queueUrl and messageBody
func NewSendMessageInput(input interface{}, queueUrl string) (*SendMessageInput, error) {

	if reflect.DeepEqual(reflect.TypeOf(input).Kind(), reflect.Ptr) {
		return nil, errors.New(ErrNoPointerParameterAllowed)
	}

	if len(queueUrl) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	sendMsgInput := new(sqs.SendMessageInput)

	msgBody, err := marshalStructToJsonString(input)
	if err != nil {
		return nil, err
	}

	sendMsgInput = sendMsgInput.SetMessageBody(msgBody).SetQueueUrl(queueUrl)

	out := new(SendMessageInput)

	out.SendMessageInput = sendMsgInput

	return out, nil

}

// marshalStructToJsonString marshals input into a string contains its json representation
func marshalStructToJsonString(input interface{}) (string, error) {

	if reflect.DeepEqual(reflect.TypeOf(input).Kind(), reflect.Ptr) {
		return "", errors.New(ErrNoPointerParameterAllowed)
	}

	b, marshalErr := json.Marshal(input)
	if marshalErr != nil {
		return "", marshalErr
	}

	return string(b), nil

}

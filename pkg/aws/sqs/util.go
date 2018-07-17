package sqs

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/service/sqs"
	"reflect"
	"encoding/base64"
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

// NewSendMessageInput returns a new *SendMessageInput initialized with queueUrl and messageBody encoded in base64
func NewSendMessageInput(input interface{}, queueUrl string) (*SendMessageInput, error) {

	if reflect.DeepEqual(reflect.TypeOf(input).Kind(), reflect.Ptr) {
		return nil, errors.New(ErrNoPointerParameterAllowed)
	}

	if len(queueUrl) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	sendMsgInput := new(sqs.SendMessageInput)

	msgBody, err := marshalStructToJson(input)
	if err != nil {
		return nil, err
	}

	b64MsgBody := base64.StdEncoding.EncodeToString(msgBody)

	sendMsgInput = sendMsgInput.SetMessageBody(b64MsgBody).SetQueueUrl(queueUrl)

	out := new(SendMessageInput)

	out.SendMessageInput = sendMsgInput

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

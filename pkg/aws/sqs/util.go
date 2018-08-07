package sqs

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/aws/aws-sdk-go/service/sqs"
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

// NewSendMessageInput returns a new *SendMessageInput initialized with queueUrl and messageBody.
// If base64Encode = true then the messageBody will be encoded in base64
func NewSendMessageInput(input interface{}, queueUrl string, base64Encode bool) (*SendMessageInput, error) {

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

	msgStringBody := string(msgBody)

	if base64Encode {
		msgStringBody = base64.StdEncoding.EncodeToString(msgBody)
	}

	sendMsgInput = sendMsgInput.SetMessageBody(msgStringBody)
	sendMsgInput = sendMsgInput.SetQueueUrl(queueUrl)

	out := new(SendMessageInput)
	out.SendMessageInput = sendMsgInput

	return out, nil

}

// NewGetQueueUrlInput returns a new GetQueueUrlInput given a queue name
func NewGetQueueUrlInput(queueName string) (*GetQueueUrlInput, error) {

	if len(queueName) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	getQueueUrlInput := new(sqs.GetQueueUrlInput)
	getQueueUrlInput = getQueueUrlInput.SetQueueName(queueName)

	out := new(GetQueueUrlInput)
	out.GetQueueUrlInput = getQueueUrlInput

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

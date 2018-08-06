package sqs

import (
	"errors"
	"reflect"

	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/sqs"
)

// CreateQueue creates a new queue given its name
func (svc *SQS) CreateQueue(queueName string) error {

	if len(queueName) == 0 {
		return errors.New(ErrEmptyParameter)
	}

	createQueueIn, createQueueInErr := sqs.NewCreateQueueInput(queueName)
	if createQueueInErr != nil {
		return createQueueInErr
	}

	err := svc.SQSCreateQueue(createQueueIn)
	if err != nil {
		return err
	}

	return nil

}

// SendMessage sends a new message on a sns queue given an input and a valid queue url
func (svc *SQS) SendMessage(input interface{}, queueUrl string) error {

	if reflect.DeepEqual(reflect.TypeOf(input).Kind(), reflect.Ptr) {
		return errors.New(ErrNoPointerParameterAllowed)
	}

	if len(queueUrl) == 0 {
		return errors.New(ErrEmptyParameter)
	}

	sendMsgIn, sendMsgInErr := sqs.NewSendMessageInput(
		input,
		queueUrl,
	)
	if sendMsgInErr != nil {
		return sendMsgInErr
	}

	err := svc.SQSSendMessage(sendMsgIn)
	if err != nil {
		return err
	}

	return nil

}

// GetQueueAttributes returns *sqs.GetQueueAttributesOutput and error is nil if queue exists
func (svc *SQS) GetQueueAttributes(queueUrl string) (*sqs.GetQueueAttributesOutput, error) {

	if len(queueUrl) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	getQueueAttrsIn, getQueueAttrsInErr := sqs.NewGetQueueAttributesInput(queueUrl)
	if getQueueAttrsInErr != nil {
		return nil, getQueueAttrsInErr
	}

	out, err := svc.SQSGetQueueAttributes(getQueueAttrsIn)
	if err != nil {
		return nil, err
	}

	return out, nil

}

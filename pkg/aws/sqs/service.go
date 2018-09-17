package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetQueueUrlInput embeds *sqs.GetQueueUrlInput
type GetQueueUrlInput struct {
	*sqs.GetQueueUrlInput
}

// SQSCreateQueue creates an sns queue given a queue name
func (svc *SQS) SQSCreateQueue(queue string) error {

	input, inputErr := NewCreateQueueInput(queue)
	if inputErr != nil {
		return inputErr
	}

	if _, err := svc.CreateQueue(input); err != nil {
		return err
	}

	return nil

}

// SQSGetQueueAttributes returns error if queue does not exist, get queue attributes otherwise
func (svc *SQS) SQSGetQueueAttributes(queueUrl string) (*sqs.GetQueueAttributesOutput, error) {

	input, inputErr := NewGetQueueAttributesInput(queueUrl)
	if inputErr != nil {
		return nil, inputErr
	}

	out, err := svc.GetQueueAttributes(input)
	if err != nil {
		return nil, err
	}

	return out, nil

}

// SQSSendMessage sends a message on SQS
func (svc *SQS) SQSSendMessage(input interface{}, queueName string, base64Encode bool) error {

	queueUrl, queueUrlErr := svc.SQSGetQueueUrl(queueName)
	if queueUrlErr != nil {
		return queueUrlErr
	}

	sendMsgInput, sendMsgInputErr := NewSendMessageInput(
		input,
		queueUrl,
		base64Encode,
	)
	if sendMsgInputErr != nil {
		return sendMsgInputErr
	}

	if _, err := svc.SendMessage(sendMsgInput); err != nil {
		return err
	}

	return nil

}

// SQSGetQueueUrl gets a queue's url given its name
func (svc *SQS) SQSGetQueueUrl(queueUrl string) (string, error) {

	input, inputErr := NewGetQueueUrlInput(queueUrl)
	if inputErr != nil {
		return "", inputErr
	}

	out, err := svc.GetQueueUrl(input)
	if err != nil {
		return "", err
	}

	return *out.QueueUrl, nil

}

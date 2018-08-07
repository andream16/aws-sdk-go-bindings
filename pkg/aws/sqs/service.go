package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

// CreateQueueInput embeds *sqs.CreateQueueInput
type CreateQueueInput struct {
	*sqs.CreateQueueInput
}

// GetQueueAttributesInput embeds *sqs.GetQueueAttributesInput
type GetQueueAttributesInput struct {
	*sqs.GetQueueAttributesInput
}

// GetQueueAttributesOutput embeds *sqs.GetQueueAttributesOutput
type GetQueueAttributesOutput struct {
	*sqs.GetQueueAttributesOutput
}

// SendMessageInput embeds *sqs.SendMessageInput
type SendMessageInput struct {
	*sqs.SendMessageInput
}

// GetQueueUrlInput embeds *sqs.GetQueueUrlInput
type GetQueueUrlInput struct {
	*sqs.GetQueueUrlInput
}

// SQSCreateQueue creates an sns queue given a *CreateQueueInput
func (svc *SQS) SQSCreateQueue(input *CreateQueueInput) error {

	if _, err := svc.CreateQueue(input.CreateQueueInput); err != nil {
		return err
	}

	return nil

}

// SQSGetQueueAttributes returns error if queue does not exist, nil otherwise
func (svc *SQS) SQSGetQueueAttributes(input *GetQueueAttributesInput) (*GetQueueAttributesOutput, error) {

	attrs, err := svc.GetQueueAttributes(input.GetQueueAttributesInput)
	if err != nil {
		return nil, err
	}

	out := new(GetQueueAttributesOutput)
	out.GetQueueAttributesOutput = attrs

	return out, nil

}

// SQSSendMessage sends a message on SQS
func (svc *SQS) SQSSendMessage(input *SendMessageInput) error {

	if _, err := svc.SendMessage(input.SendMessageInput); err != nil {
		return err
	}

	return nil

}

// SQSGetQueueUrl gets a queue's url given its name
func (svc *SQS) SQSGetQueueUrl(input *GetQueueUrlInput) (string, error) {

	out, err := svc.GetQueueUrl(input.GetQueueUrlInput)
	if err != nil {
		return "", nil
	}

	return *out.QueueUrl, nil

}

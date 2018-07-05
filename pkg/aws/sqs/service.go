package sqs

import "github.com/aws/aws-sdk-go/service/sqs"

// CreateQueueInput embeds *sqs.CreateQueueInput
type CreateQueueInput struct {
	*sqs.CreateQueueInput
}

// GetQueueAttributesInput embeds *sqs.GetQueueAttributesInput
type GetQueueAttributesInput struct {
	*sqs.GetQueueAttributesInput
}

// SendMessageInput embeds *sqs.SendMessageInput
type SendMessageInput struct {
	*sqs.SendMessageInput
}

// SQSCreateQueue creates an sns queue given a *CreateQueueInput
func (svc *SQS) SQSCreateQueue(input *CreateQueueInput) error {

	if _, err := svc.CreateQueue(input.CreateQueueInput); err != nil {
		return err
	}

	return nil

}

// SQSGetQueueAttributes returns error if queue does not exist, nil otherwise
func (svc *SQS) SQSGetQueueAttributes(input *GetQueueAttributesInput) error {

	if _, err := svc.GetQueueAttributes(input.GetQueueAttributesInput); err != nil {
		return err
	}

	return nil

}

// SQSSendMessage sends a message on SQS
func (svc *SQS) SQSSendMessage(input *SendMessageInput) error {

	if _, err := svc.SendMessage(input.SendMessageInput); err != nil {
		return err
	}

	return nil

}

package sqs

import (
	"encoding/base64"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/pkg/errors"
)

// SQSer describes sqs API.
type SQSer interface {
	CreateQueue(name string) error
	GetQueueAttributes(url string) (*sqs.GetQueueAttributesOutput, error)
	GetQueueURL(name string) (string, error)
	SendMessage(payload []byte, url string, encBase64 bool) error
}

// SQS is the alias for SQS.
type SQS struct {
	sqs sqsiface.SQSAPI
}

// New returns a new SQS.
func New(config *aws.Config) (*SQS, error) {

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize sns session")
	}

	return &SQS{
		sqs: sqs.New(sess),
	}, nil

}

// CreateQueue creates a new queue.
func (s SQS) CreateQueue(name string) error {

	if name == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "name")
	}

	in := &sqs.CreateQueueInput{
		QueueName: aws.String(name),
	}

	_, err := s.sqs.CreateQueue(in)
	if err != nil {
		return errors.Wrapf(err, "unable to create queue %s", name)
	}

	return nil

}

// GetQueueAttributes gets queue attributes.
func (s SQS) GetQueueAttributes(url string) (*sqs.GetQueueAttributesOutput, error) {

	if url == "" {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "url")
	}

	in := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(url),
	}

	attrs, err := s.sqs.GetQueueAttributes(in)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get attributes for queue with url %s", url)
	}

	return attrs, nil

}

// GetQueueURL gets queue's url given its name.
func (s SQS) GetQueueURL(name string) (string, error) {

	if name == "" {
		return "", errors.Wrap(bindings.ErrInvalidParameter, "name")
	}

	in := &sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	}

	out, err := s.sqs.GetQueueUrl(in)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get url for queue with name %s", name)
	}

	return *out.QueueUrl, nil

}

// SendMessage sends a payload in a queue.
// If encBase64 is true, the payload will be encoded in base64.
func (s SQS) SendMessage(payload []byte, url string, encBase64 bool) error {

	if len(payload) == 0 {
		return errors.Wrap(bindings.ErrInvalidParameter, "payload")
	}
	if url == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "url")
	}

	b := string(payload)
	if encBase64 {
		b = base64.StdEncoding.EncodeToString(payload)
	}

	in := &sqs.SendMessageInput{
		MessageBody: aws.String(b),
		QueueUrl:    aws.String(url),
	}

	_, err := s.sqs.SendMessage(in)
	if err != nil {
		return errors.Wrapf(err, "unable to send message on queue %s", url)
	}

	return nil

}

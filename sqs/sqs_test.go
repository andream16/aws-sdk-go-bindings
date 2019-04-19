package sqs

import (
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/pkg/errors"
)

type mockSQSClient struct {
	sqsiface.SQSAPI
}

func (*mockSQSClient) CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	return nil, nil
}

func (*mockSQSClient) GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	return nil, nil
}

func (*mockSQSClient) GetQueueUrl(*sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return &sqs.GetQueueUrlOutput{
		QueueUrl: aws.String("someUrl"),
	}, nil
}

func (*mockSQSClient) SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return nil, nil
}

type mockFailingSQSClient struct {
	sqsiface.SQSAPI
}

func (*mockFailingSQSClient) CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	return nil, errors.New("some error")
}

func (*mockFailingSQSClient) GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	return nil, errors.New("some error")
}

func (*mockFailingSQSClient) GetQueueUrl(*sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return nil, errors.New("some error")
}

func (*mockFailingSQSClient) SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return nil, errors.New("some error")
}

func TestSQS_CreateQueue(t *testing.T) {

	t.Run("should return an error because name is empty", func(t *testing.T) {

		s := &SQS{}

		err := s.CreateQueue("")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected invalid name error, got %s", err)
		}

	})

	t.Run("should return an error because CreateQueue failed", func(t *testing.T) {

		mockSvc := &mockFailingSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		err := s.CreateQueue("someName")
		if err == nil {
			t.Fatal("expected error from create queue, got nil")
		}

	})

	t.Run("should successfully create a queue", func(t *testing.T) {

		mockSvc := &mockSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		err := s.CreateQueue("someName")
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

func TestSQS_GetQueueAttributes(t *testing.T) {

	t.Run("should return an error because url is empty", func(t *testing.T) {

		s := &SQS{}

		_, err := s.GetQueueAttributes("")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected invalid url error, got %s", err)
		}

	})

	t.Run("should return an error because GetQueueAttributes failed", func(t *testing.T) {

		mockSvc := &mockFailingSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		_, err := s.GetQueueAttributes("someUrl")
		if err == nil {
			t.Fatal("expected error from get queue attributes, got nil")
		}

	})

	t.Run("should successfully get queue attributes", func(t *testing.T) {

		mockSvc := &mockSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		_, err := s.GetQueueAttributes("someUrl")
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

func TestSQS_GetQueueURL(t *testing.T) {

	t.Run("should return an error because name is empty", func(t *testing.T) {

		s := &SQS{}

		_, err := s.GetQueueURL("")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected invalid name error, got %s", err)
		}

	})

	t.Run("should return an error because GetQueueUrl failed", func(t *testing.T) {

		mockSvc := &mockFailingSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		_, err := s.GetQueueURL("someName")
		if err == nil {
			t.Fatal("expected error from get queue url, got nil")
		}

	})

	t.Run("should successfully get queue url", func(t *testing.T) {

		mockSvc := &mockSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		url, err := s.GetQueueURL("someName")
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

		if url == "" {
			t.Fatal("expected some url, got an empty value")
		}

	})

}

func TestSQS_SendMessage(t *testing.T) {

	t.Run("should return an error because payload is empty", func(t *testing.T) {

		s := &SQS{}

		err := s.SendMessage(nil, "", false)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected invalid payload error, got %s", err)
		}

	})

	t.Run("should return an error because url is empty", func(t *testing.T) {

		s := &SQS{}

		err := s.SendMessage([]byte{1, 121, 31}, "", false)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected invalid url error, got %s", err)
		}

	})

	t.Run("should return an error because GetQueueUrl failed", func(t *testing.T) {

		mockSvc := &mockFailingSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		err := s.SendMessage([]byte{1, 121, 31}, "someUrl", true)
		if err == nil {
			t.Fatal("expected error from send message, got nil")
		}

	})

	t.Run("should successfully get queue url", func(t *testing.T) {

		mockSvc := &mockSQSClient{}

		s := &SQS{
			sqs: mockSvc,
		}

		err := s.SendMessage([]byte{1, 121, 31}, "someUrl", false)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

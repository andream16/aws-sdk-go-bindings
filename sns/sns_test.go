package sns

import (
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/pkg/errors"
)

type mockSNSClient struct {
	snsiface.SNSAPI
}

func (*mockSNSClient) Publish(*sns.PublishInput) (*sns.PublishOutput, error) {
	return nil, nil
}

type mockFailingSNSClient struct {
	snsiface.SNSAPI
}

func (*mockFailingSNSClient) Publish(*sns.PublishInput) (*sns.PublishOutput, error) {
	return nil, errors.New("some error")
}

func TestSNS_Publish(t *testing.T) {

	t.Run("should return an error because payload is empty", func(t *testing.T) {

		s := &SNS{}

		err := s.Publish(nil, "", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error for invalid target, got %s", err)
		}

	})

	t.Run("should return an error because target is empty", func(t *testing.T) {

		s := &SNS{}

		err := s.Publish([]byte{1, 121, 3}, "", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error for invalid target, got %s", err)
		}

	})

	t.Run("should return an error because Publish failed", func(t *testing.T) {

		mockSvc := &mockFailingSNSClient{}

		s := &SNS{
			sns: mockSvc,
		}

		err := s.Publish([]byte{40}, "someTarget", "")
		if err == nil {
			t.Fatal("expected publish error, got nil")
		}

	})

	t.Run("should successfully publish a non json message", func(t *testing.T) {

		mockSvc := &mockSNSClient{}

		s := &SNS{
			sns: mockSvc,
		}

		err := s.Publish([]byte{40}, "someTarget", "someStructure")
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

package s3

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pkg/errors"

	bindings "github.com/andream16/aws-sdk-go-bindings"
)

type mockS3Client struct {
	s3iface.S3API
}

func (m *mockS3Client) CreateBucket(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return nil, nil
}

func (m *mockS3Client) GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	r := bytes.NewReader([]byte{10, 21, 121})
	return &s3.GetObjectOutput{
		Body: ioutil.NopCloser(r),
	}, nil
}

type mockFailingS3Client struct {
	s3iface.S3API
}

func (m *mockFailingS3Client) CreateBucket(in *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return nil, errors.New("some error")
}

func (m *mockFailingS3Client) GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return nil, errors.New("some error")
}

func TestNew(t *testing.T) {

	t.Run("should return an error because region is missing", func(t *testing.T) {
		_, err := New("")
		if err == nil {
			t.Fatal("expected missing required parameter region error, got nil")
		}
	})

}

func TestCreateBucket(t *testing.T) {

	t.Run("should return an error because the bucket is empty", func(t *testing.T) {

		s, err := New("eu-central-1")

		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

		err = s.CreateBucket("")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because there was an error creating the bucket", func(t *testing.T) {

		mockSvc := &mockFailingS3Client{}

		s := S3{
			s3: mockSvc,
		}

		err := s.CreateBucket("some bucket")
		if err == nil {
			t.Fatal("expected create bucket error, got nil")
		}

	})

	t.Run("should successfully create a bucket", func(t *testing.T) {

		mockSvc := &mockS3Client{}

		s := S3{
			s3: mockSvc,
		}

		err := s.CreateBucket("some bucket")
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

func TestGetObject(t *testing.T) {

	t.Run("should return an error because the bucket is empty", func(t *testing.T) {

		s, err := New("eu-central-1")

		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

		_, err = s.GetObject("", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because path is empty", func(t *testing.T) {

		s, err := New("eu-central-1")

		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

		_, err = s.GetObject("someBucket", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because there was an error getting the object", func(t *testing.T) {

		mockSvc := &mockFailingS3Client{}

		s := S3{
			s3: mockSvc,
		}

		_, err := s.GetObject("someBucket", "somePath")
		if err == nil {
			t.Fatal("expected get object error, got nil")
		}

	})

	t.Run("should successfully get an object", func(t *testing.T) {

		mockSvc := &mockS3Client{}

		s := S3{
			s3: mockSvc,
		}

		_, err := s.GetObject("someBucket", "somePath")
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

package s3

import (
	"bytes"
	"io/ioutil"
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pkg/errors"
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

func (m *mockS3Client) PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return nil, nil
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

func (m *mockFailingS3Client) PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return nil, errors.New("some error")
}

func TestCreateBucket(t *testing.T) {

	t.Run("should return an error because bucket is empty", func(t *testing.T) {

		s := &S3{}

		err := s.CreateBucket("")
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

	t.Run("should return an error because bucket is empty", func(t *testing.T) {

		s := &S3{}

		_, err := s.GetObject("", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because path is empty", func(t *testing.T) {

		s := &S3{}

		_, err := s.GetObject("someBucket", "")
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

func TestPutObject(t *testing.T) {

	t.Run("should return an error because bucket is empty", func(t *testing.T) {

		s := &S3{}

		err := s.PutObject("", "", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because object path is empty", func(t *testing.T) {

		s := &S3{}

		err := s.PutObject("someBucket", "", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because object name is empty", func(t *testing.T) {

		s := &S3{}

		err := s.PutObject("someBucket", "somePath", "")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because some error happened during read file", func(t *testing.T) {

		s := &S3{}

		err := s.PutObject("someBucket", "somePath", "someName")
		if err == nil {
			t.Fatal("expected read file error, got nil")
		}

	})

	t.Run("should return an error because some error happened put object", func(t *testing.T) {

		mockSvc := &mockFailingS3Client{}

		s := S3{
			s3: mockSvc,
		}

		err := s.PutObject("someBucket", "testdata/putobjecttest.jpg", "someName")
		if err == nil {
			t.Fatal("expected put object error, got nil")
		}

	})

	t.Run("should successfully put an object", func(t *testing.T) {

		mockSvc := &mockS3Client{}

		s := S3{
			s3: mockSvc,
		}

		err := s.PutObject("someBucket", "testdata/putobjecttest.jpg", "someName")
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

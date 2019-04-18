package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pkg/errors"

	bindings "github.com/andream16/aws-sdk-go-bindings"
	"github.com/andream16/aws-sdk-go-bindings/internal/format"
)

// S3er describes s3 API.
type S3er interface {
	CreateBucket(string) error
	GetObject(bucket, path string) ([]byte, error)
	PutObject(bucket, objectName, objectPath string)
}

// S3 is the alias for s3.
type S3 struct {
	s3 s3iface.S3API
}

// New returns a new S3.
func New(region string, options ...bindings.Option) (*S3, error) {

	if region == "" {
		return nil, errors.New("required parameter region cannot be empty")
	}

	cfg := &bindings.Config{
		Region: format.StrToPtr(region),
	}

	for _, option := range options {
		option(cfg)
	}

	sess, err := session.NewSession((*aws.Config)(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize s3 session")
	}

	return &S3{
		s3: s3.New(sess),
	}, nil

}

// CreateBucket creates an s3 bucket.
func (s S3) CreateBucket(bucket string) error {

	if bucket == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "bucket")
	}

	in := &s3.CreateBucketInput{
		Bucket: format.StrToPtr(bucket),
	}

	_, err := s.s3.CreateBucket(in)
	if err != nil {
		return errors.Wrapf(err, "unable to create bucket %s", bucket)
	}

	return nil

}

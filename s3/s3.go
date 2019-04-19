package s3

import (
	"bytes"
	"io/ioutil"
	"net/http"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pkg/errors"
)

// S3er describes s3 API.
type S3er interface {
	CreateBucket(string) error
	GetObject(bucket, path string) ([]byte, error)
	PutObject(bucket, path, name string) error
}

// S3 is the alias for s3.
type S3 struct {
	s3 s3iface.S3API
}

type file struct {
	content       []byte
	contentType   string
	contentLength int64
}

// New returns a new S3.
func New(region string, options ...bindings.Option) (*S3, error) {

	if region == "" {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "region cannot be empty")
	}

	cfg := &bindings.Config{
		Region: aws.String(region),
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
		Bucket: aws.String(bucket),
	}

	_, err := s.s3.CreateBucket(in)
	if err != nil {
		return errors.Wrapf(err, "unable to create bucket %s", bucket)
	}

	return nil

}

// GetObject fetches an object from a given path.
func (s S3) GetObject(bucket, path string) ([]byte, error) {

	if bucket == "" {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "bucket")
	}
	if path == "" {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "path")
	}

	in := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}

	out, err := s.s3.GetObject(in)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get object bucket %s/%s", bucket, path)
	}

	return ioutil.ReadAll(out.Body)

}

// PutObject uploads an object to a bucket.
func (s S3) PutObject(bucket, path, name string) error {

	if bucket == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "bucket")
	}
	if path == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "path")
	}
	if name == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "name")
	}

	file, err := readFile(path)
	if err != nil {
		return errors.Wrapf(err, "coudln't read local file from path %s", path)
	}

	in := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(name),
		ContentType:   aws.String(file.contentType),
		Body:          bytes.NewReader(file.content),
		ContentLength: aws.Int64(file.contentLength),
	}

	_, err = s.s3.PutObject(in)
	if err != nil {
		return errors.Wrapf(err, "unable to put object %s", name)
	}

	return nil

}

func readFile(path string) (*file, error) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	contentLength := int64(len(b))
	contentType := http.DetectContentType(b)

	return &file{
		content:       b,
		contentType:   contentType,
		contentLength: contentLength,
	}, nil

}

package s3

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

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
	PutObject(bucket, objectName, objectPath string) error
}

// S3 is the alias for s3.
type S3 struct {
	s3 s3iface.S3API
}

type file struct {
	body          []byte
	contentType   string
	contentLength int64
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

// GetObject fetches an object from a given path.
func (s S3) GetObject(bucket, path string) ([]byte, error) {

	if bucket == "" {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "bucket")
	}
	if path == "" {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "path")
	}

	in := &s3.GetObjectInput{
		Bucket: format.StrToPtr(bucket),
		Key:    format.StrToPtr(path),
	}

	out, err := s.s3.GetObject(in)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get object bucket %s", bucket+"/"+path)
	}

	return ioutil.ReadAll(out.Body)

}

// PutObject uploads an object to a bucket.
func (s S3) PutObject(bucket, objectName, objectPath string) error {

	if bucket == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "bucket")
	}
	if objectName == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "object name")
	}
	if objectPath == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "object path")
	}

	file, err := readFile(objectPath)
	if err != nil {
		return errors.Wrapf(err, "coudln't read local file from path %s", objectPath)
	}

	in := &s3.PutObjectInput{
		Bucket:        format.StrToPtr(bucket),
		Key:           format.StrToPtr(objectName),
		ContentType:   format.StrToPtr(file.contentType),
		Body:          bytes.NewReader(file.body),
		ContentLength: format.Int64ToPtr(file.contentLength),
	}

	_, err = s.s3.PutObject(in)
	if err != nil {
		return errors.Wrapf(err, "unable to put object %s", objectName)
	}

	return nil

}

func readFile(path string) (*file, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	contentLength := fileInfo.Size()
	buffer := make([]byte, contentLength)

	_, err = f.Read(buffer)
	if err != nil {
		return nil, err
	}

	contentType := http.DetectContentType(buffer)

	return &file{
		body:          buffer,
		contentType:   contentType,
		contentLength: contentLength,
	}, nil

}

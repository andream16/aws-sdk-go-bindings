package s3

import (
	"bytes"
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
)

// NewCreateBucketInput returns a new *CreateBucketInput
func NewCreateBucketInput(bucketName string) (*CreateBucketInput, error) {

	if len(bucketName) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	createBucketInput := new(s3.CreateBucketInput)
	createBucketInput = createBucketInput.SetBucket(bucketName)

	out := new(CreateBucketInput)
	out.CreateBucketInput = createBucketInput

	return out, nil

}

// NewGetObjectInput returns a new *GetObjectInput given a bucket and a source image
func NewGetObjectInput(bucket, source string) (*GetObjectInput, error) {

	if len(bucket) == 0 || len(source) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	getObjectInput := new(s3.GetObjectInput)
	getObjectInput = getObjectInput.SetBucket(bucket)
	getObjectInput = getObjectInput.SetKey(source)

	out := new(GetObjectInput)
	out.GetObjectInput = getObjectInput

	return out, nil

}

// NewPutObjectInput returns a new *PutObjectInput
func NewPutObjectInput(bucket, fileName, contentType string, image []byte, size int64) (*PutObjectInput, error) {

	if len(bucket) == 0 || len(fileName) == 0 || len(contentType) == 0 || len(image) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	putObjectInput := new(s3.PutObjectInput)
	putObjectInput = putObjectInput.SetBucket(bucket)
	putObjectInput = putObjectInput.SetKey(fileName)
	putObjectInput = putObjectInput.SetContentType(contentType)
	putObjectInput = putObjectInput.SetBody(bytes.NewReader(image))
	putObjectInput = putObjectInput.SetContentLength(size)

	out := new(PutObjectInput)
	out.PutObjectInput = putObjectInput

	return out, nil

}

// UnmarshalIOReadCloser extracts []byte from input.Body
func UnmarshalIOReadCloser(input io.ReadCloser) ([]byte, error) {

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(input)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

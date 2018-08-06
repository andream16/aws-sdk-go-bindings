package s3

import (
	"bytes"
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewGetObjectInput returns a new *GetObjectInput given a bucket and a source image
func NewGetObjectInput(bucket, source string) (*GetObjectInput, error) {

	if len(bucket) == 0 || len(source) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	out := new(GetObjectInput)
	out.GetObjectInput = &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(source),
	}

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

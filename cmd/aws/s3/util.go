package s3

import (
	"bytes"
	"errors"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
	"io/ioutil"
)

// UnmarshalGetObjectOutput extracts bytes from *s3.GetObjectOutput
func UnmarshalGetObjectOutput(input *s3.GetObjectOutput) ([]byte, error) {

	if *input.ContentLength == 0 {
		return nil, errors.New(ErrEmptyContentLength)
	}

	body, bytesErr := ioutil.ReadAll(input.Body)
	if bytesErr != nil {
		return nil, bytesErr
	}
	if len(body) == 0 {
		return nil, errors.New(ErrEmptyBody)
	}

	input.Body = ioutil.NopCloser(bytes.NewReader(body))

	b, err := s3.UnmarshalIOReadCloser(input.Body)
	if err != nil {
		return nil, err
	}

	return b, nil

}

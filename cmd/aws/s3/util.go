package s3

import "github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"

// UnmarshalGetObjectOutput extracts bytes from *s3.GetObjectOutput
func UnmarshalGetObjectOutput(input *s3.GetObjectOutput) ([]byte, error) {

	b, err := s3.UnmarshalIOReadCloser(input.Body)

	if err != nil {
		return nil, err
	}

	return b, nil

}

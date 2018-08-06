package s3

import (
	"errors"

	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
)

// GetObject returns a []byte encoding the image
func (svc *S3) GetObject(bucket, sourceImage string) ([]byte, error) {

	if len(bucket) == 0 || len(sourceImage) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	s3In, s3InErr := s3.NewGetObjectInput(
		bucket,
		sourceImage,
	)
	if s3InErr != nil {
		return nil, s3InErr
	}

	obj, objErr := svc.S3GetObject(s3In)
	if objErr != nil {
		return nil, objErr
	}

	out, err := UnmarshalGetObjectOutput(obj)
	if err != nil {
		return nil, err
	}

	return out, nil

}

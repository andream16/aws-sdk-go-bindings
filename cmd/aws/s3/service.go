package s3

import "github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"

// GetObject returns a []byte encoding the image
func (svc *S3) GetObject(bucket, sourceImage string) ([]byte, error) {

	s3In := s3.NewGetObjectInput(
		bucket,
		sourceImage,
	)

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

package s3

import (
	"errors"

	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
)

// CreateBucket creates an S3 bucket given a bucket name
func (svc *S3) CreateBucket(bucket string) error {

	if len(bucket) == 0 {
		return errors.New(ErrEmptyParameter)
	}

	in, inErr := s3.NewCreateBucketInput(bucket)
	if inErr != nil {
		return inErr
	}

	err := svc.S3CreateBucket(in)
	if err != nil {
		return err
	}

	return nil

}

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

// PutObject adds a given object on S3
func (svc *S3) PutObject(bucket, objectName, objectPath string) error {

	if len(bucket) == 0 || len(objectName) == 0 || len(objectPath) == 0 {
		return errors.New(ErrEmptyParameter)
	}

	imgMeta, readErr := ReadImage(objectPath)
	if readErr != nil {
		return readErr
	}

	in, inErr := s3.NewPutObjectInput(
		bucket,
		objectName,
		imgMeta.ContentType,
		imgMeta.Body,
		imgMeta.ContentSize,
	)
	if inErr != nil {
		return inErr
	}

	err := svc.S3PutObject(in)
	if err != nil {
		return err
	}

	return nil

}

package s3

import "github.com/aws/aws-sdk-go/service/s3"

// CreateBucketInput embeds *s3.CreateBucketInput
type CreateBucketInput struct {
	*s3.CreateBucketInput
}

// GetObjectInput embeds *s3.GetObjectInput
type GetObjectInput struct {
	*s3.GetObjectInput
}

// GetObjectOutput embeds *s3.GetObjectOutput
type GetObjectOutput struct {
	*s3.GetObjectOutput
}

// PutObjectInput embeds *s3.PutObjectInput
type PutObjectInput struct {
	*s3.PutObjectInput
}

// S3CreateBucket creates a new bucket given a
func (svc *S3) S3CreateBucket(input *CreateBucketInput) error {

	_, err := svc.S3.CreateBucket(input.CreateBucketInput)
	if err != nil {
		return err
	}

	return nil
}

// S3GetObject retrieves an object from S3 given a *GetObjectInput
func (svc *S3) S3GetObject(input *GetObjectInput) (*GetObjectOutput, error) {

	getObjectOut, err := svc.S3.GetObject(input.GetObjectInput)
	if err != nil {
		return nil, err
	}

	out := new(GetObjectOutput)
	out.GetObjectOutput = getObjectOut

	return out, nil

}

// S3PutObject puts a given Object on S3
func (svc *S3) S3PutObject(input *PutObjectInput) error {

	_, err := svc.S3.PutObject(input.PutObjectInput)
	if err != nil {
		return err
	}

	return nil

}

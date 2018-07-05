package s3

import "github.com/aws/aws-sdk-go/service/s3"

// GetObjectInput embeds *s3.GetObjectInput
type GetObjectInput struct {
	*s3.GetObjectInput
}

// GetObjectOutput embeds *s3.GetObjectOutput
type GetObjectOutput struct {
	*s3.GetObjectOutput
}

// S3GetObject retrieves an object from S3 given a GetObjectInput
func (svc *S3) S3GetObject(input *GetObjectInput) (*GetObjectOutput, error) {

	getObjectOut, err := svc.S3.GetObject(input.GetObjectInput)
	if err != nil {
		return nil, err
	}

	out := new(GetObjectOutput)
	out.GetObjectOutput = getObjectOut

	return out, nil

}

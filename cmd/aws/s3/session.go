package s3

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
)

// S3 embeds *s3.S3 and is used to call sns methods on high level
type S3 struct {
	*s3.S3
}

// New returns a *S3 given a *aws.Session
func New(svc *aws.Session) (*S3, error) {

	s3Svc, err := s3.New(svc.Session)
	if err != nil {
		return nil, err
	}

	newS3Svc := new(S3)
	newS3Svc.S3 = s3Svc

	return newS3Svc, nil

}

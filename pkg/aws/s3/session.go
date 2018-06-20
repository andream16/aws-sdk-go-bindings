package s3

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 embeds *s3.S3 to be used to call New
type S3 struct {
	*s3.S3
}

// New returns a new *S3 embedding *s3.S3
func New(svc *aws.Session) (*S3, error) {

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	s3Svc := new(S3)
	s3Svc.S3 = s3.New(newSvc)

	return s3Svc, nil

}

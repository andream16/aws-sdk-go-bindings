package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
)

// S3 embeds *s3.S3 to be used to call New
type S3 struct {
	*s3.S3
}

// New returns a new *S3 embedding *s3.S3
func New(svc *pkgAws.Session, endpoint string) (*S3, error) {

	if len(endpoint) > 0 {
		svc.Config.Endpoint = aws.String(endpoint)
	}

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	s3Svc := &S3{
		S3: s3.New(newSvc),
	}

	return s3Svc, nil

}

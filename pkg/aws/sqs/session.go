package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
)

// SQS embeds *sns.SNS
type SQS struct {
	*sqs.SQS
}

// New returns a new *SQS
func New(svc *pkgAws.Session, endpoint string) (*SQS, error) {

	if endpoint != "" {
		svc.Config.Endpoint = aws.String(endpoint)
	}

	newSvc, err := session.NewSession(svc.Config)
	if err != nil {
		return nil, err
	}

	sqsSvc := &SQS{
		SQS: sqs.New(newSvc),
	}

	return sqsSvc, nil

}

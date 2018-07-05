package sqs

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQS embeds *sns.SNS
type SQS struct {
	*sqs.SQS
}

// New returns a new *SQS
func New(svc *aws.Session) (*SQS, error) {

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	sqsSvc := new(SQS)
	sqsSvc.SQS = sqs.New(newSvc)

	return sqsSvc, nil

}

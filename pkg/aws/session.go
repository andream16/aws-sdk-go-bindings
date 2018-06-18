package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Session struct {
	*session.Session
}

// New returns a new *Session embedding *session.Session
func New() (*Session, error) {

	cfg := new(aws.Config)

	awsSession, awsSessionErr := session.NewSession(cfg)
	if awsSessionErr != nil {
		return nil, awsSessionErr
	}

	svc := new(Session)

	svc.Session = awsSession

	return svc, nil

}

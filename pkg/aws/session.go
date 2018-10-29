package aws

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Session embeds session.Session
type Session struct {
	*session.Session
}

// New returns a new *Session embedding *session.Session
func New(input *SessionInput) (*Session, error) {

	if input.region == "" {
		return nil, errors.New(ErrNoRegionProvided)
	}

	cfg := &aws.Config{
		Region: aws.String(input.region),
	}

	awsSession, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	svc := &Session{}

	svc.Session = awsSession

	return svc, nil

}

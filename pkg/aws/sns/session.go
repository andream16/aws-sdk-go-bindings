package sns

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// Session embeds sns.SNS to be used to call New
type Session struct {
	*sns.SNS
}

// New returns a new *Session embedding *sns.SNS
func New(svc *session.Session) (*Session, error) {

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	snsSvc := new(Session)
	snsSvc.SNS = sns.New(newSvc)

	return snsSvc, nil

}

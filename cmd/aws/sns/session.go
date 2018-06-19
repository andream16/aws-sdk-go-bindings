package sns

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/sns"
)

// New returns a *sns.Session given a *aws.Session
func New(svc *aws.Session) (snsSvc *sns.Session, err error) {

	snsSvc, err = sns.New(svc.Session)
	return

}

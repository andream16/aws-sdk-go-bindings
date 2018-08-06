package sns

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
)

// SNS embeds sns.SNS to be used to call New
type SNS struct {
	*sns.SNS
}

// New returns a new *SNS embedding *sns.SNS
func New(svc *aws.Session) (*SNS, error) {

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	snsSvc := new(SNS)
	snsSvc.SNS = sns.New(newSvc)

	return snsSvc, nil

}

package sns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
)

// SNS embeds sns.SNS to be used to call New
type SNS struct {
	*sns.SNS
}

// New returns a new *SNS embedding *sns.SNS
func New(svc *pkgAws.Session, endpoint string) (*SNS, error) {

	if len(endpoint) > 0 {
		svc.Config.Endpoint = aws.String(endpoint)
	}

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	snsSvc := new(SNS)
	snsSvc.SNS = sns.New(newSvc)

	return snsSvc, nil

}

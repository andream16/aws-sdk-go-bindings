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

	if endpoint != "" {
		svc.Config.Endpoint = aws.String(endpoint)
	}

	newSvc, err := session.NewSession(svc.Config)
	if err != nil {
		return nil, err
	}

	snsSvc := &SNS{
		SNS: sns.New(newSvc),
	}

	return snsSvc, nil

}

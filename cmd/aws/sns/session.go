package sns

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/sns"
)

// SNS embeds *sns.SNS and is used to call sns methods on high level
type SNS struct {
	*sns.SNS
}

// New returns a *SNS given a *aws.Session and an optional endpoint
func New(svc *aws.Session, endpoint string) (*SNS, error) {

	snsSvc, err := sns.New(svc.Session, endpoint)
	if err != nil {
		return nil, err
	}

	newSnsSvc := new(SNS)
	newSnsSvc.SNS = snsSvc

	return newSnsSvc, nil

}

package aws

import "github.com/andream16/aws-sdk-go-bindings/pkg/aws"

// Session embeds a *aws.Session
type Session struct {
	*aws.Session
}

// New returns a new *aws.Session
func New(region string) (*Session, error) {

	in := aws.NewSessionInput(region)

	svc, err := aws.New(in)
	if err != nil {
		return nil, err
	}

	newSvc := new(Session)
	newSvc.Session = svc

	return newSvc, nil

}

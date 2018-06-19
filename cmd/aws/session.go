package aws

import "github.com/andream16/aws-sdk-go-bindings/pkg/aws"

// New returns a new *aws.Session
func New(region string) (svc *aws.Session, err error) {

	in := aws.NewSessionInput(region)

	svc, err = aws.New(in)

	return

}

package sns

import "github.com/aws/aws-sdk-go/service/sns"

type PublishInput struct {
	*sns.PublishInput
}

// SnsPublish publishes a *PublishInput on SNS
func (svc *Session) SnsPublish(in *PublishInput) (err error) {

	_, err = svc.SNS.Publish(in.PublishInput)
	return

}

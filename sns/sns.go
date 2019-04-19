package sns

import (
	"encoding/json"
	"strings"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/pkg/errors"
)

// SNSer describes sns API.
type SNSer interface {
	Publish(payload []byte, target string, msgStr string) error
}

// SNS is the alias for SNS.
type SNS struct {
	sns snsiface.SNSAPI
}

type message struct {
	Default string `json:"default"`
}

// New returns a new S3.
func New(config *aws.Config) (*SNS, error) {

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize sns session")
	}

	return &SNS{
		sns: sns.New(sess),
	}, nil

}

// Publish publishes a payload in a given target arn.
// If msgStr is not passed, a default `json` structure will be used.
func (s SNS) Publish(payload []byte, target string, msgStr string) error {

	if len(payload) == 0 {
		return errors.Wrap(bindings.ErrInvalidParameter, "payload")
	}
	if target == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "target")
	}

	in, err := newPublishInput(payload, target, msgStr)
	if err != nil {
		return errors.Wrap(err, "unable to build valid publish input")
	}

	_, err = s.sns.Publish(in)
	if err != nil {
		return errors.Wrapf(err, "unable to publish payload on target %s", target)
	}

	return nil

}

func newPublishInput(payload []byte, target, msgStr string) (*sns.PublishInput, error) {

	if msgStr != "" && msgStr != "json" {
		return &sns.PublishInput{
			Message:          aws.String(string(payload)),
			MessageStructure: aws.String(msgStr),
			TargetArn:        aws.String(target),
		}, nil
	}

	// Mandatory since SNS needs escaped `"`. So we need to escape them to `\"`
	unquote := strings.Replace(string(payload), `"`, "\"", -1)

	// Mandatory since SNS needs a payload like:
	// {
	// 		"default" : {
	// 			\"par1\" : \"some value\"
	// 		}
	// }
	snsBody := message{
		Default: unquote,
	}

	// Mandatory since we want to get a string out of encoded bytes
	b, err := json.Marshal(snsBody)
	if err != nil {
		return nil, err
	}

	return &sns.PublishInput{
		Message:          aws.String(string(b)),
		MessageStructure: aws.String("json"),
		TargetArn:        aws.String(target),
	}, nil

}

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
	Publish(payload []byte, target string, messageStructure string) error
}

// SNS is the alias for SNS.
type SNS struct {
	sns snsiface.SNSAPI
}

type message struct {
	Default string `json:"default"`
}

// New returns a new S3.
func New(region string, options ...bindings.Option) (*SNS, error) {

	if region == "" {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "region cannot be empty")
	}

	cfg := &bindings.Config{
		Region: aws.String(region),
	}

	for _, option := range options {
		option(cfg)
	}

	sess, err := session.NewSession((*aws.Config)(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize sns session")
	}

	return &SNS{
		sns: sns.New(sess),
	}, nil

}

// Publish publishes a payload in a given target arn.
// If messageStructure is not passed, a default `json` structure will be used.
func (s SNS) Publish(payload []byte, target string, messageStructure string) error {

	if target == "" {
		return errors.Wrap(bindings.ErrInvalidParameter, "target")
	}

	msgStructure := "json"
	if messageStructure != "" {
		msgStructure = messageStructure
	}

	in, err := newPublishInput(payload, target, msgStructure)
	if err != nil {
		return errors.Wrap(err, "unable to build valid publish input")
	}

	_, err = s.sns.Publish(in)
	if err != nil {
		return errors.Wrapf(err, "unable to publish payload on target %s", target)

	}

	return nil

}

func newPublishInput(payload []byte, target, messageStructure string) (*sns.PublishInput, error) {

	// Mandatory since SNS needs escaped `"`. So we need to escape them to `\"`
	unquote := strings.Replace(string(payload), `"`, "\"", -1)

	// Mandatory since SNS needs bodies like:
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
		MessageStructure: aws.String(messageStructure),
		TargetArn:        aws.String(target),
	}, nil

}

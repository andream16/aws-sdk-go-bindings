package bindings

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"

	"github.com/andream16/aws-sdk-go-bindings/internal/format"
)

// Session is the alias for aws session.
type Session session.Session

// Config is the alias for aws config.
type Config aws.Config

// Option is used to set session's options.
type Option func(cfg *Config)

// New returns a new Session.
func New(region string, options ...func(cfg *Config)) (*Session, error) {

	if region == "" {
		return nil, errors.New("required parameter region cannot be empty")
	}

	cfg := &Config{
		Region: format.StrToPtr(region),
	}

	for _, option := range options {
		option(cfg)
	}

	sess, err := session.NewSession((*aws.Config)(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize session")
	}

	return (*Session)(sess), nil

}

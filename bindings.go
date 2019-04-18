package bindings

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
)

// Config is the alias for aws config.
type Config aws.Config

// Option is used to set session's options.
type Option func(cfg *Config)

// Package level related errors.
var (
	ErrInvalidParameter = errors.New("invalid_parameter")
)

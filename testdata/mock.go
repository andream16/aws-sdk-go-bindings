package testdata

import (
	"testing"
	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/stretchr/testify/assert"
)

func MockConfiguration(t *testing.T) configuration.Configuration {

	t.Helper()

	cfg, err := configuration.Get()

	assert.NoError(t, err)
	assert.NotEmpty(t, cfg)

	return cfg

}

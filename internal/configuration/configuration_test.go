package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	cfg, err := Get()

	assert.NoError(t, err)
	assert.NotEmpty(t, cfg)
	assert.NotEqual(t, 0, len(cfg.Region))

}

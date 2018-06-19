package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {

	cfg, err := Get()

	assert.NoError(t, err)
	assert.NotEmpty(t, cfg)
	assert.NotEqual(t, 0, len(cfg.Region))

}

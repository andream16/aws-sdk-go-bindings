package aws

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	svc, err := New()

	assert.NotEmpty(t, svc)
	assert.NoError(t, err)

}

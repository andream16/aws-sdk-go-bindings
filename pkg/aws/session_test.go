package aws

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	region := "eu-central-1"

	in := NewSessionInput(region)

	assert.NotEmpty(t, in)

	svc, err := New(in)

	assert.NotEmpty(t, svc)
	assert.NoError(t, err)

}

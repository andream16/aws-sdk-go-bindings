package aws

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	in := NewSessionInput(cfg.Region)

	assert.NotEmpty(t, in)

	svc, err := New(in)

	assert.NotEmpty(t, svc)
	assert.NoError(t, err)

}

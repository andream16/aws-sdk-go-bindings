package aws

import (
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc, err := New(cfg.Region)

	assert.NoError(t, err)
	assert.NotEmpty(t, svc)

}

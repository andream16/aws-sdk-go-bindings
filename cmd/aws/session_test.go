package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc, err := New(cfg.Region)

	assert.NoError(t, err)
	assert.NotEmpty(t, svc)

}

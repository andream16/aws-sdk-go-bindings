package s3

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc, svcErr := aws.New(cfg.Region)

	assert.NoError(t, svcErr)

	s3Svc, s3SvcErr := New(svc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

}

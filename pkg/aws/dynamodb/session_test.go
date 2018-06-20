package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	in := aws.NewSessionInput(cfg.Region)

	awsSvc, awsSvcErr := aws.New(in)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	snsSvc, snsSvcErr := New(awsSvc, "")

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

}

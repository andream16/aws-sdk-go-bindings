package rekognition

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	in, inErr := aws.NewSessionInput(cfg.Region)

	assert.NoError(t, inErr)

	awsSvc, awsSvcErr := aws.New(in)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	snsSvc, snsSvcErr := New(awsSvc, cfg.Region)

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

}

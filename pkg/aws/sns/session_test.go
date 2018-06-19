package sns

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	region := "eu-central-1"

	in := aws.NewSessionInput(region)

	awsSvc, awsSvcErr := aws.New(in)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	snsSvc, snsSvcErr := New(awsSvc.Session)

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

}

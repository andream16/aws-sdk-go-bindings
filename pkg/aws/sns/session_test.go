package sns

import (
	"testing"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	awsSvc, awsSvcErr := aws.New()

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	snsSvc, snsSvcErr := New(awsSvc.Session)

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

}

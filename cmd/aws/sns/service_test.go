package sns

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestPublishType struct {
	SomeParam string `json:"some_param"`
}

func TestSns_Publish(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	svc, svcErr := New(awsSvc.Session)

	assert.NoError(t, svcErr)
	assert.NotEmpty(t, svc)

	err := svc.Publish(
		TestPublishType{"some_val"},
		cfg.SNS.TargetArn,
	)

	assert.NoError(t, err)

}

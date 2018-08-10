package sns

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestSession_SnsPublish(t *testing.T) {

	body := `{"default":"{\"par1\":\"pr1\",\"par2\":\"pr2\"}"}`

	cfg := testdata.MockConfiguration(t)

	svcIn, svcInErr := pkgAws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)
	awsSvc, awsSvcErr := pkgAws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	snsSvc, snsSvcErr := New(awsSvc, cfg.SNS.Endpoint)

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

	err := snsSvc.SnsPublish(
		body,
		cfg.SNS.TargetArn,
	)

	assert.NoError(t, err)

	shouldBeEmptyParameterErr := snsSvc.SnsPublish(
		body,
		"",
	)

	assert.Error(t, shouldBeEmptyParameterErr)
	assert.Contains(t, shouldBeEmptyParameterErr.Error(), ErrEmptyParameter)

	shouldBeNoPointerParameterAllowedErr := snsSvc.SnsPublish(
		&body,
		cfg.SNS.TargetArn,
	)

	assert.Error(t, shouldBeNoPointerParameterAllowedErr)
	assert.Contains(t, shouldBeNoPointerParameterAllowedErr.Error(), ErrPointerParameterNotAllowed)

}

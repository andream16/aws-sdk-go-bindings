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

	svcIn, err := pkgAws.NewSessionInput(cfg.Region)

	assert.NoError(t, err)
	awsSvc, err := pkgAws.New(svcIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, awsSvc)

	snsSvc, err := New(awsSvc, cfg.SNS.Endpoint)

	assert.NoError(t, err)
	assert.NotEmpty(t, snsSvc)

	err = snsSvc.SnsPublish(
		body,
		cfg.SNS.TargetArn,
	)

	assert.NoError(t, err)

	err = snsSvc.SnsPublish(
		body,
		"",
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

	err = snsSvc.SnsPublish(
		&body,
		cfg.SNS.TargetArn,
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrPointerParameterNotAllowed)

}

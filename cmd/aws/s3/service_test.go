package s3

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestS3_GetObject(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	svc, svcErr := New(awsSvc)

	assert.NoError(t, svcErr)
	assert.NotEmpty(t, svc)

	b, err := svc.GetObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
	)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(b))

}

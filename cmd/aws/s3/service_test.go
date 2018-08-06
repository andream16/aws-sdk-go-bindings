package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
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

	_, shouldBeErr1 := svc.GetObject(
		"",
		cfg.S3.SourceImage,
	)

	assert.Error(t, shouldBeErr1)
	assert.Equal(t, ErrEmptyParameter, shouldBeErr1.Error())

	_, shouldBeErr2 := svc.GetObject(
		"",
		cfg.S3.SourceImage,
	)

	assert.Error(t, shouldBeErr2)
	assert.Equal(t, ErrEmptyParameter, shouldBeErr2.Error())

}

package s3

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestS3_S3GetObject(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svcIn := aws.NewSessionInput(cfg.Region)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	in := NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
	)

	out, err := s3Svc.S3GetObject(in)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

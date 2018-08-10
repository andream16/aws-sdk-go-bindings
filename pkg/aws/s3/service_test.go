package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestS3_S3CreateBucket(t *testing.T) {

	s3Svc, cfg := newS3Svc(t)

	createBucket(cfg, s3Svc, t)

}

func TestS3_S3GetObject(t *testing.T) {

	s3Svc, cfg := newS3Svc(t)

	createBucket(cfg, s3Svc, t)
	putObject(t)

	out, err := s3Svc.S3GetObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func TestS3_S3PutObject(t *testing.T) {

	putObject(t)

}

func createBucket(cfg *configuration.Configuration, svc *S3, t *testing.T) {

	svc.S3CreateBucket(cfg.S3.Bucket)

}

func newS3Svc(t *testing.T) (*S3, *configuration.Configuration) {

	t.Helper()

	cfg := testdata.MockConfiguration(t)

	svcIn, svcInErr := aws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := New(awsSvc, cfg.S3.Endpoint)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	return s3Svc, cfg

}

func putObject(t *testing.T) {

	s3Svc, cfg := newS3Svc(t)

	createBucket(cfg, s3Svc, t)

	err := s3Svc.S3PutObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
		"../../../assets/compare_faces_test-source.jpg",
	)

	assert.NoError(t, err)

}

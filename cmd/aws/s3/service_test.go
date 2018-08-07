package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestS3_CreateBucket(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := mockS3Svc(cfg, t)

	err := svc.CreateBucket(cfg.S3.Bucket)

	if err != nil {
		assert.Contains(t, err.Error(), "BucketAlreadyExists")
	} else {
		assert.NoError(t, err)
	}

}

func TestS3_GetObject(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := mockS3Svc(cfg, t)

	svc.CreateBucket(cfg.S3.Bucket)

	putObjErr := svc.PutObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
		"../../../assets/compare_faces_test-source.jpg",
	)

	assert.NoError(t, putObjErr)

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

func TestS3_PutObject(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc := mockS3Svc(cfg, t)

	err := svc.PutObject(
		cfg.S3.Bucket,
		"some_name",
		"../../../assets/compare_faces_test-source.jpg",
	)

	assert.NoError(t, err)

}

func mockS3Svc(cfg *configuration.Configuration, t *testing.T) *S3 {

	t.Helper()

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	svc, svcErr := New(awsSvc, cfg.S3.Endpoint)

	assert.NoError(t, svcErr)
	assert.NotEmpty(t, svc)

	return svc

}

package rekognition

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws/s3"
	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRekognition_Methods(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	testRekognitionCompareFaces(t, cfg)
	testRekognitionDetectFaces(t, cfg)
	testRekognitionDetectText(t, cfg)

}

func testRekognitionCompareFaces(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	encodedSourceObject, encodedSourceObjectErr := s3Svc.GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.SourceImage,
	)

	assert.NoError(t, encodedSourceObjectErr)
	assert.NotEqual(t, 0, len(encodedSourceObject))

	encodedTargetObject, encodedTargetObjectErr := s3Svc.GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.TargetImage,
	)

	assert.NoError(t, encodedTargetObjectErr)
	assert.NotEqual(t, 0, len(encodedTargetObject))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	out, err := rekSvc.CompareFaces(
		encodedSourceObject,
		encodedTargetObject,
		cfg.Rekognition.CompareFaces.Similarity,
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func testRekognitionDetectFaces(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	encodedObject, encodedObjectErr := s3Svc.GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectFaces.SourceImage,
	)

	assert.NoError(t, encodedObjectErr)
	assert.NotEqual(t, 0, len(encodedObject))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	out, err := rekSvc.DetectFaces(encodedObject)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func testRekognitionDetectText(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	encodedObject, encodedObjectErr := s3Svc.GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectText.SourceImage,
	)

	assert.NoError(t, encodedObjectErr)
	assert.NotEqual(t, 0, len(encodedObject))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	out, err := rekSvc.DetectText(encodedObject)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

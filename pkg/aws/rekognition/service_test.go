package rekognition

import (
	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRekognition_RekognitionMethods(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	testRekognitionRekognitionCompareFaces(t, cfg)
	testRekognitionRekognitionDetectFaces(t, cfg)
	testRekognitionRekognitionDetectText(t, cfg)

}

func testRekognitionRekognitionCompareFaces(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	svcIn, svcInErr := aws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	getSourceObjectIn, getSourceObjectInErr := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.SourceImage,
	)

	assert.NoError(t, getSourceObjectInErr)

	getSourceObjectOut, getSourceObjectErr := s3Svc.S3GetObject(getSourceObjectIn)

	assert.NoError(t, getSourceObjectErr)

	bSource, bSourceErr := s3.UnmarshalIOReadCloser(getSourceObjectOut.Body)

	assert.NoError(t, bSourceErr)
	assert.NotEqual(t, 0, len(bSource))

	getTargetObjectIn, getTargetObjectInErr := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.TargetImage,
	)

	assert.NoError(t, getTargetObjectInErr)

	getTargetObjectOut, getTargetObjectOutErr := s3Svc.S3GetObject(getTargetObjectIn)

	assert.NoError(t, getTargetObjectOutErr)

	bTarget, bTargetErr := s3.UnmarshalIOReadCloser(getTargetObjectOut.Body)

	assert.NoError(t, bTargetErr)
	assert.NotEqual(t, 0, len(bTarget))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	rekIn, rekInErr := NewCompareFacesInput(
		bSource,
		bTarget,
		cfg.Rekognition.CompareFaces.Similarity,
	)

	assert.NoError(t, rekInErr)

	out, err := rekSvc.RekognitionCompareFaces(rekIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func testRekognitionRekognitionDetectFaces(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	svcIn, svcInErr := aws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	getObjectIn, getObjectInErr := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectFaces.SourceImage,
	)

	assert.NoError(t, getObjectInErr)

	getObjectOut, getObjectErr := s3Svc.S3GetObject(getObjectIn)

	assert.NoError(t, getObjectErr)

	b, bErr := s3.UnmarshalIOReadCloser(getObjectOut.Body)

	assert.NoError(t, bErr)
	assert.NotEqual(t, 0, len(b))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	rekIn, rekInErr := NewDetectFacesInput(b)

	assert.NoError(t, rekInErr)

	out, err := rekSvc.RekognitionDetectFaces(rekIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func testRekognitionRekognitionDetectText(t *testing.T, cfg configuration.Configuration) {

	t.Helper()

	svcIn, svcInErr := aws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	getObjectIn, getObjectInErr := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectText.SourceImage,
	)

	assert.NoError(t, getObjectInErr)

	getObjectOut, getObjectErr := s3Svc.S3GetObject(getObjectIn)

	assert.NoError(t, getObjectErr)

	b, bErr := s3.UnmarshalIOReadCloser(getObjectOut.Body)

	assert.NoError(t, bErr)
	assert.NotEqual(t, 0, len(b))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	rekIn, rekInErr := NewDetectTextInput(b)

	assert.NoError(t, rekInErr)

	out, err := rekSvc.RekognitionDetectText(rekIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

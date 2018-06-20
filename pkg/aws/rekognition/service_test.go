package rekognition

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRekognition_RekognitionCompareFaces(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svcIn := aws.NewSessionInput(cfg.Region)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	getSourceObjectIn := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.SourceImage,
	)

	getSourceObjectOut, getSourceObjectErr := s3Svc.S3GetObject(getSourceObjectIn)

	assert.NoError(t, getSourceObjectErr)

	bSource, bSourceErr := s3.UnmarshalIOReadCloser(getSourceObjectOut.Body)

	assert.NoError(t, bSourceErr)
	assert.NotEqual(t, 0, len(bSource))

	getTargetObjectIn := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.TargetImage,
	)

	getTargetObjectOut, getTargetObjectOutErr := s3Svc.S3GetObject(getTargetObjectIn)

	assert.NoError(t, getTargetObjectOutErr)

	bTarget, bTargetErr := s3.UnmarshalIOReadCloser(getTargetObjectOut.Body)

	assert.NoError(t, bTargetErr)
	assert.NotEqual(t, 0, len(bTarget))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	rekIn := NewCompareFacesInput(
		bSource,
		bTarget,
		cfg.Rekognition.CompareFaces.Similarity,
	)

	out, err := rekSvc.RekognitionCompareFaces(rekIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func TestRekognition_RekognitionDetectFaces(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svcIn := aws.NewSessionInput(cfg.Region)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	getObjectIn := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectFaces.SourceImage,
	)

	getObjectOut, getObjectErr := s3Svc.S3GetObject(getObjectIn)

	assert.NoError(t, getObjectErr)

	b, bErr := s3.UnmarshalIOReadCloser(getObjectOut.Body)

	assert.NoError(t, bErr)
	assert.NotEqual(t, 0, len(b))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	rekIn := NewDetectFacesInput(b)

	out, err := rekSvc.RekognitionDetectFaces(rekIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func TestRekognition_RekognitionDetectText(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svcIn := aws.NewSessionInput(cfg.Region)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	getObjectIn := s3.NewGetObjectInput(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectText.SourceImage,
	)

	getObjectOut, getObjectErr := s3Svc.S3GetObject(getObjectIn)

	assert.NoError(t, getObjectErr)

	b, bErr := s3.UnmarshalIOReadCloser(getObjectOut.Body)

	assert.NoError(t, bErr)
	assert.NotEqual(t, 0, len(b))

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	rekIn := NewDetectTextInput(b)

	out, err := rekSvc.RekognitionDetectText(rekIn)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

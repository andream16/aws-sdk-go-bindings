package rekognition

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws/s3"
	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestRekognition_Methods(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	s3Svc, rekSvc := mockSessions(cfg, t)

	uploadImages(cfg, s3Svc, t)

	funcs := []func(*configuration.Configuration, *s3.S3, *Rekognition, *testing.T){
		testRekognitionCompareFaces,
		testRekognitionDetectFaces,
		testRekognitionDetectText,
	}

	for i := 0; i < len(funcs); i++ {

		funcs[i](cfg, s3Svc, rekSvc, t)

	}

}

func testRekognitionCompareFaces(cfg *configuration.Configuration, s3Svc *s3.S3, rekSvc *Rekognition, t *testing.T) {

	t.Helper()

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

	out, err := rekSvc.CompareFaces(
		encodedSourceObject,
		encodedTargetObject,
		cfg.Rekognition.CompareFaces.Similarity,
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	_, shouldBeErr1 := rekSvc.CompareFaces(
		[]byte{},
		encodedTargetObject,
		cfg.Rekognition.CompareFaces.Similarity,
	)

	_, shouldBeErr2 := rekSvc.CompareFaces(
		encodedSourceObject,
		[]byte{},
		cfg.Rekognition.CompareFaces.Similarity,
	)

	assert.Error(t, shouldBeErr1, shouldBeErr2)
	assert.Equal(t, ErrEmptyBytes, shouldBeErr1.Error(), shouldBeErr2.Error())

	_, shouldBeErr3 := rekSvc.CompareFaces(
		encodedSourceObject,
		encodedTargetObject,
		0,
	)

	assert.Error(t, shouldBeErr3)
	assert.Equal(t, ErrBadSimilarity, shouldBeErr3.Error())

}

func testRekognitionDetectFaces(cfg *configuration.Configuration, s3Svc *s3.S3, rekSvc *Rekognition, t *testing.T) {

	t.Helper()

	encodedObject, encodedObjectErr := s3Svc.GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectFaces.SourceImage,
	)

	assert.NoError(t, encodedObjectErr)
	assert.NotEqual(t, 0, len(encodedObject))

	out, err := rekSvc.DetectFaces(encodedObject)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	_, shouldBeErr := rekSvc.DetectFaces([]byte{})

	assert.Error(t, shouldBeErr)
	assert.Equal(t, ErrEmptyBytes, shouldBeErr.Error())

}

func testRekognitionDetectText(cfg *configuration.Configuration, s3Svc *s3.S3, rekSvc *Rekognition, t *testing.T) {

	t.Helper()

	encodedObject, encodedObjectErr := s3Svc.GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectText.SourceImage,
	)

	assert.NoError(t, encodedObjectErr)
	assert.NotEqual(t, 0, len(encodedObject))

	out, err := rekSvc.DetectText(encodedObject)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	_, shouldBeErr := rekSvc.DetectText([]byte{})

	assert.Error(t, shouldBeErr)
	assert.Equal(t, ErrEmptyBytes, shouldBeErr.Error())

}

func uploadImages(cfg *configuration.Configuration, svc *s3.S3, t *testing.T) {

	t.Helper()

	images := map[string]string{
		cfg.Rekognition.CompareFaces.SourceImage: "../../../assets/compare_faces_test-source.jpg",
		cfg.Rekognition.CompareFaces.TargetImage: "../../../assets/compare_faces_test-target.jpg",
		cfg.Rekognition.DetectFaces.SourceImage:  "../../../assets/detect_faces_test-source.jpg",
		cfg.Rekognition.DetectText.SourceImage:   "../../../assets/detect_text_test-source.jpg",
	}

	for k, v := range images {

		k := k
		v := v

		go func() {

			err := svc.PutObject(cfg.S3.Bucket, k, v)

			assert.NoError(t, err)

		}()

	}

	time.Sleep(100 * time.Millisecond)

}

func mockSessions(cfg *configuration.Configuration, t *testing.T) (*s3.S3, *Rekognition) {

	t.Helper()

	awsSvc, awsSvcErr := aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc, cfg.S3.Endpoint)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	awsSvc, awsSvcErr = aws.New(cfg.Region)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	return s3Svc, rekSvc

}

package rekognition

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestRekognition_RekognitionMethods(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	s3Svc, rekSvc := mockSessions(cfg, t)

	createBucketInput, createBucketInputErr := s3.NewCreateBucketInput(cfg.S3.Bucket)

	assert.NoError(t, createBucketInputErr)
	assert.NotEmpty(t, createBucketInput)

	s3Svc.S3CreateBucket(cfg.S3.Bucket)

	uploadImages(cfg, s3Svc, t)

	funcs := []func(*configuration.Configuration, *s3.S3, *Rekognition, *testing.T){
		testRekognitionRekognitionCompareFaces,
		testRekognitionRekognitionDetectFaces,
		testRekognitionRekognitionDetectText,
	}

	for i := 0; i < len(funcs); i++ {
		funcs[i](cfg, s3Svc, rekSvc, t)
	}

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

			err := svc.S3PutObject(
				cfg.S3.Bucket,
				k,
				v,
			)

			assert.NoError(t, err)

		}()

	}

	time.Sleep(100 * time.Millisecond)

}

func readImage(path string, t *testing.T) ([]byte, string, int64) {

	file, fileErr := os.Open(path)

	assert.NoError(t, fileErr)
	assert.NotEmpty(t, file)

	defer file.Close()

	fileInfo, fileInfoErr := file.Stat()

	assert.NoError(t, fileInfoErr)

	size := fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)
	fileType := http.DetectContentType(buffer)

	return buffer, fileType, size

}

func testRekognitionRekognitionCompareFaces(cfg *configuration.Configuration, s3Svc *s3.S3, rekSvc *Rekognition, t *testing.T) {

	t.Helper()

	getSourceObjectOut, getSourceObjectErr := s3Svc.S3GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.SourceImage,
	)

	assert.NoError(t, getSourceObjectErr)
	assert.NotEqual(t, 0, len(getSourceObjectOut))

	getTargetObjectOut, getTargetObjectOutErr := s3Svc.S3GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.CompareFaces.TargetImage,
	)

	assert.NoError(t, getTargetObjectOutErr)
	assert.NotEqual(t, 0, len(getTargetObjectOut))

	out, err := rekSvc.RekognitionCompareFaces(
		getSourceObjectOut,
		getTargetObjectOut,
		cfg.Rekognition.CompareFaces.Similarity,
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func testRekognitionRekognitionDetectFaces(cfg *configuration.Configuration, s3Svc *s3.S3, rekSvc *Rekognition, t *testing.T) {

	t.Helper()

	getObjectOut, getObjectErr := s3Svc.S3GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectFaces.SourceImage,
	)

	assert.NoError(t, getObjectErr)
	assert.NotEqual(t, 0, len(getObjectOut))

	out, err := rekSvc.RekognitionDetectFaces(getObjectOut)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func testRekognitionRekognitionDetectText(cfg *configuration.Configuration, s3Svc *s3.S3, rekSvc *Rekognition, t *testing.T) {

	t.Helper()

	getObjectOut, getObjectErr := s3Svc.S3GetObject(
		cfg.S3.Bucket,
		cfg.Rekognition.DetectText.SourceImage,
	)

	assert.NoError(t, getObjectErr)
	assert.NotEqual(t, 0, len(getObjectOut))

	out, err := rekSvc.RekognitionDetectText(getObjectOut)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func mockSessions(cfg *configuration.Configuration, t *testing.T) (*s3.S3, *Rekognition) {

	t.Helper()

	svcIn, svcInErr := aws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)

	awsSvc, awsSvcErr := aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	s3Svc, s3SvcErr := s3.New(awsSvc, cfg.S3.Endpoint)

	assert.NoError(t, s3SvcErr)
	assert.NotEmpty(t, s3Svc)

	awsSvc, awsSvcErr = aws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	rekSvc, rekSvcErr := New(awsSvc, cfg.Rekognition.Region)

	assert.NoError(t, rekSvcErr)
	assert.NotEmpty(t, rekSvc)

	return s3Svc, rekSvc

}

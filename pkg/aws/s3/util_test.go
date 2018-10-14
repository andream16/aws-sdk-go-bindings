package s3

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestReadImageOutput_SetBody(t *testing.T) {

	body := []byte("some_body")

	readImageOutput := &ReadImageOutput{}
	readImageOutput = readImageOutput.SetBody(body)

	assert.Equal(t, body, readImageOutput.Body)

}

func TestReadImageOutput_SetContentType(t *testing.T) {

	contentType := "some_content_type"

	readImageOutput := &ReadImageOutput{}
	readImageOutput = readImageOutput.SetContentType(contentType)

	assert.Equal(t, contentType, readImageOutput.ContentType)

}

func TestReadImageOutput_SetContentSize(t *testing.T) {

	var contentSize int64 = 10

	readImageOutput := &ReadImageOutput{}
	readImageOutput = readImageOutput.SetContentSize(contentSize)

	assert.Equal(t, contentSize, readImageOutput.ContentSize)

}

func TestUnmarshalGetObjectOutput(t *testing.T) {

	s := "create a really cool md5 checksum of me"
	body := []byte(s)

	var getObjectOutputMock = &s3.GetObjectOutput{
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: aws.Int64(int64(len(body))),
	}

	out, err := UnmarshalGetObjectOutput(getObjectOutputMock)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	_, errEmptyContentLength := UnmarshalGetObjectOutput(
		&s3.GetObjectOutput{
			Body:          ioutil.NopCloser(bytes.NewReader(body)),
			ContentLength: aws.Int64(0),
		},
	)

	assert.Error(t, errEmptyContentLength)
	assert.Contains(t, errEmptyContentLength.Error(), ErrEmptyContentLength)

	_, errEmptyBody := UnmarshalGetObjectOutput(
		&s3.GetObjectOutput{
			Body:          ioutil.NopCloser(bytes.NewReader([]byte{})),
			ContentLength: aws.Int64(40),
		},
	)

	assert.Error(t, errEmptyBody)
	assert.Contains(t, errEmptyBody.Error(), ErrEmptyBody)

}

func TestReadImage(t *testing.T) {

	imgPath := "../../../assets/compare_faces_test-source.jpg"

	out, err := ReadImage(imgPath)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func TestNewCreateBucketInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewCreateBucketInput(cfg.S3.Bucket)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	_, shouldBeEmptyErr := NewCreateBucketInput("")

	assert.Error(t, shouldBeEmptyErr)
	assert.Contains(t, shouldBeEmptyErr.Error(), ErrEmptyParameter)

}

func TestNewGetObjectInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetObjectInput(cfg.S3.Bucket, cfg.S3.SourceImage)

	assert.NoError(t, err)
	assert.Equal(t, cfg.S3.Bucket, *out.Bucket)
	assert.Equal(t, cfg.S3.SourceImage, *out.Key)

	_, shouldBeEmptyErr1 := NewGetObjectInput("", cfg.S3.SourceImage)
	assert.Contains(t, shouldBeEmptyErr1.Error(), ErrEmptyParameter)
	_, shouldBeEmptyErr2 := NewGetObjectInput(cfg.S3.Bucket, "")
	assert.Contains(t, shouldBeEmptyErr2.Error(), ErrEmptyParameter)

}

func TestNewPutObjectInput(t *testing.T) {

	contentType := "jpg"
	var contentSize int64 = 16

	cfg := testdata.MockConfiguration(t)

	putObjectInput, putObjectInputErr := NewPutObjectInput(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
		contentType,
		[]byte(cfg.S3.SourceImage),
		16,
	)

	assert.NoError(t, putObjectInputErr)
	assert.NotEmpty(t, putObjectInput)

	_, shouldBeEmptyErr1 := NewPutObjectInput("", cfg.S3.SourceImage, contentType, []byte(cfg.S3.SourceImage), contentSize)
	assert.Contains(t, shouldBeEmptyErr1.Error(), ErrEmptyParameter)
	_, shouldBeEmptyErr2 := NewPutObjectInput(cfg.S3.Bucket, "", contentType, []byte(cfg.S3.SourceImage), contentSize)
	assert.Contains(t, shouldBeEmptyErr2.Error(), ErrEmptyParameter)
	_, shouldBeEmptyErr3 := NewPutObjectInput(cfg.S3.Bucket, cfg.S3.SourceImage, "", []byte(cfg.S3.SourceImage), contentSize)
	assert.Contains(t, shouldBeEmptyErr3.Error(), ErrEmptyParameter)
	_, shouldBeEmptyErr4 := NewPutObjectInput(cfg.S3.Bucket, cfg.S3.SourceImage, contentType, []byte(""), contentSize)
	assert.Contains(t, shouldBeEmptyErr4.Error(), ErrEmptyParameter)

}

func TestUnmarshalIOReadCloser(t *testing.T) {

	body := []byte("something")

	readCloser := ioutil.NopCloser(bytes.NewReader(body))

	out, err := UnmarshalIOReadCloser(readCloser)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

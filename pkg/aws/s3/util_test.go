package s3

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestNewCreateBucketInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewCreateBucketInput(cfg.S3.Bucket)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	_, shouldBeEmptyErr := NewCreateBucketInput("")

	assert.Error(t, shouldBeEmptyErr)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr.Error())

}

func TestNewGetObjectInput(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	out, err := NewGetObjectInput(cfg.S3.Bucket, cfg.S3.SourceImage)

	assert.NoError(t, err)
	assert.Equal(t, cfg.S3.Bucket, *out.Bucket)
	assert.Equal(t, cfg.S3.SourceImage, *out.Key)

	_, shouldBeEmptyErr1 := NewGetObjectInput("", cfg.S3.SourceImage)
	_, shouldBeEmptyErr2 := NewGetObjectInput(cfg.S3.Bucket, "")
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr1.Error(), shouldBeEmptyErr2.Error())

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
	_, shouldBeEmptyErr2 := NewPutObjectInput(cfg.S3.Bucket, "", contentType, []byte(cfg.S3.SourceImage), contentSize)
	_, shouldBeEmptyErr3 := NewPutObjectInput(cfg.S3.Bucket, cfg.S3.SourceImage, "", []byte(cfg.S3.SourceImage), contentSize)
	_, shouldBeEmptyErr4 := NewPutObjectInput(cfg.S3.Bucket, cfg.S3.SourceImage, contentType, []byte(""), contentSize)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr1.Error(), shouldBeEmptyErr2.Error(), shouldBeEmptyErr3.Error(), shouldBeEmptyErr4.Error())

}

func TestUnmarshalIOReadCloser(t *testing.T) {

	body := []byte("something")

	readCloser := ioutil.NopCloser(bytes.NewReader(body))

	out, err := UnmarshalIOReadCloser(readCloser)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

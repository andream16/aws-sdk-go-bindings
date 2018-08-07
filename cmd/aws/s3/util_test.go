package s3

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"

	pkgS3 "github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
)

func TestReadImageOutput_SetBody(t *testing.T) {

	body := []byte("some_body")

	readImageOutput := new(ReadImageOutput)
	readImageOutput = readImageOutput.SetBody(body)

	assert.Equal(t, body, readImageOutput.Body)

}

func TestReadImageOutput_SetContentType(t *testing.T) {

	contentType := "some_content_type"

	readImageOutput := new(ReadImageOutput)
	readImageOutput = readImageOutput.SetContentType(contentType)

	assert.Equal(t, contentType, readImageOutput.ContentType)

}

func TestReadImageOutput_SetContentSize(t *testing.T) {

	var contentSize int64 = 10

	readImageOutput := new(ReadImageOutput)
	readImageOutput = readImageOutput.SetContentSize(contentSize)

	assert.Equal(t, contentSize, readImageOutput.ContentSize)

}

func TestUnmarshalGetObjectOutput(t *testing.T) {

	s := "create a really cool md5 checksum of me"
	body := []byte(s)

	var getObjectOutputMock = &pkgS3.GetObjectOutput{
		GetObjectOutput: &s3.GetObjectOutput{
			Body:          ioutil.NopCloser(bytes.NewReader(body)),
			ContentLength: aws.Int64(int64(len(body))),
		},
	}

	out, err := UnmarshalGetObjectOutput(getObjectOutputMock)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	_, errEmptyContentLength := UnmarshalGetObjectOutput(
		&pkgS3.GetObjectOutput{
			GetObjectOutput: &s3.GetObjectOutput{
				Body:          ioutil.NopCloser(bytes.NewReader(body)),
				ContentLength: aws.Int64(0),
			},
		},
	)

	assert.Error(t, errEmptyContentLength)
	assert.Equal(t, ErrEmptyContentLength, errEmptyContentLength.Error())

	_, errEmptyBody := UnmarshalGetObjectOutput(
		&pkgS3.GetObjectOutput{
			GetObjectOutput: &s3.GetObjectOutput{
				Body:          ioutil.NopCloser(bytes.NewReader([]byte{})),
				ContentLength: aws.Int64(40),
			},
		},
	)

	assert.Error(t, errEmptyBody)
	assert.Equal(t, ErrEmptyBody, errEmptyBody.Error())

}

func TestReadImage(t *testing.T) {

	imgPath := "../../../assets/compare_faces_test-source.jpg"

	out, err := ReadImage(imgPath)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

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

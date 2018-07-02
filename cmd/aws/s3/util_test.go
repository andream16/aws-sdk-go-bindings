package s3

import (
	"bytes"
	pkgS3 "github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
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

}

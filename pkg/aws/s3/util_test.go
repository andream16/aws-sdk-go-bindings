package s3

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewGetObjectInput(t *testing.T) {

	bucket := "some_bucket"
	img := "some_img"

	out, err := NewGetObjectInput(bucket, img)

	assert.NoError(t, err)
	assert.Equal(t, bucket, *out.Bucket)
	assert.Equal(t, img, *out.Key)

	_, shouldBeEmptyErr1 := NewGetObjectInput("", img)
	_, shouldBeEmptyErr2 := NewGetObjectInput(bucket, "")
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr1.Error(), shouldBeEmptyErr2.Error())

}

func TestUnmarshalIOReadCloser(t *testing.T) {

	body := []byte("something")

	readCloser := ioutil.NopCloser(bytes.NewReader(body))

	out, err := UnmarshalIOReadCloser(readCloser)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

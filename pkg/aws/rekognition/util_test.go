package rekognition

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCompareFacesInput(t *testing.T) {

	target := []byte("someTarget")
	source := []byte("someSource")
	similarity := 90.0

	in := NewCompareFacesInput(source, target, similarity)

	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.SourceImage.Bytes)
	assert.Equal(t, target, in.TargetImage.Bytes)
	assert.Equal(t, similarity, *in.SimilarityThreshold)

}

func TestNewDetectFacesInput(t *testing.T) {

	source := []byte("someSource")

	in := NewDetectFacesInput(source)

	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.Image.Bytes)

}

func TestNewDetectTextInput(t *testing.T) {

	source := []byte("someSource")

	in := NewDetectTextInput(source)

	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.Image.Bytes)

}

func TestNewInputImage(t *testing.T) {

	source := []byte("someSource")

	img := newInputImage(source)

	assert.NotEmpty(t, img)
	assert.Equal(t, source, img.Bytes)

}

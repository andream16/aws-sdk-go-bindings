package rekognition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCompareFacesInput(t *testing.T) {

	target := []byte("someTarget")
	source := []byte("someSource")
	similarity := 90.0

	in, inErr := NewCompareFacesInput(source, target, similarity)

	assert.NoError(t, inErr)
	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.SourceImage.Bytes)
	assert.Equal(t, target, in.TargetImage.Bytes)
	assert.Equal(t, similarity, *in.SimilarityThreshold)

	_, shouldBeEmptyErr1 := NewCompareFacesInput([]byte{}, target, similarity)
	_, shouldBeEmptyErr2 := NewCompareFacesInput(source, []byte{}, similarity)

	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr1.Error(), shouldBeEmptyErr2.Error())

	_, shouldBeBadSimilarityErr := NewCompareFacesInput(source, target, 0.0)

	assert.Equal(t, ErrBadSimilarityParameter, shouldBeBadSimilarityErr.Error())

}

func TestNewDetectFacesInput(t *testing.T) {

	source := []byte("someSource")

	in, inErr := NewDetectFacesInput(source)

	assert.NoError(t, inErr)

	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.Image.Bytes)

}

func TestNewDetectTextInput(t *testing.T) {

	source := []byte("someSource")

	in, inErr := NewDetectTextInput(source)

	assert.NoError(t, inErr)

	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.Image.Bytes)

}

func TestNewInputImage(t *testing.T) {

	source := []byte("someSource")

	img, imgErr := newInputImage(source)

	assert.NoError(t, imgErr)
	assert.NotEmpty(t, img)
	assert.Equal(t, source, img.Bytes)

}

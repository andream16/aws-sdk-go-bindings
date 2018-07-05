package rekognition

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// NewCompareFacesInput builds a CompareFacesInput starting from the two images, their bucket and a similarity threshold
func NewCompareFacesInput(source, target []byte, similarity float64) (*CompareFacesInput, error) {

	if len(source) == 0 || len(target) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	if similarity == 0 {
		return nil, errors.New(ErrBadSimilarityParameter)
	}

	newSourceInputImg, newSourceInputImgErr := newInputImage(source)
	if newSourceInputImgErr != nil {
		return nil, newSourceInputImgErr
	}
	newTargetInputImg, newTargetInputImgErr := newInputImage(target)
	if newTargetInputImgErr != nil {
		return nil, newTargetInputImgErr
	}

	out := new(CompareFacesInput)

	out.CompareFacesInput = &rekognition.CompareFacesInput{
		SimilarityThreshold: aws.Float64(similarity),
		SourceImage:         newSourceInputImg,
		TargetImage:         newTargetInputImg,
	}

	return out, nil

}

// NewDetectFacesInput builds a DetectFacesInput starting from the image
func NewDetectFacesInput(source []byte) (*DetectFacesInput, error) {

	if len(source) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	newInputImg, newInputImgErr := newInputImage(source)
	if newInputImgErr != nil {
		return nil, newInputImgErr
	}

	out := new(DetectFacesInput)

	out.DetectFacesInput = &rekognition.DetectFacesInput{
		Image: newInputImg,
	}

	return out, nil

}

// NewDetectTextInput builds a DetectTextInput starting from the image
func NewDetectTextInput(source []byte) (*DetectTextInput, error) {

	if len(source) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	newInputImg, newInputImgErr := newInputImage(source)
	if newInputImgErr != nil {
		return nil, newInputImgErr
	}

	out := new(DetectTextInput)

	out.DetectTextInput = &rekognition.DetectTextInput{
		Image: newInputImg,
	}

	return out, nil

}

// newInputImage returns a *rekognition.Image given an S3 image []byte encoded
func newInputImage(image []byte) (*rekognition.Image, error) {

	if len(image) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	out := new(rekognition.Image)

	out.Bytes = image

	return out, nil

}

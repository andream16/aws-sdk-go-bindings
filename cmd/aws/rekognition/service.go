package rekognition

import (
	"errors"

	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/rekognition"
)

// CompareFaces returns a *CompareFacesOutput
func (svc *Rekognition) CompareFaces(sourceImage, targetImage []byte, similarity float64) (*CompareFacesOutput, error) {

	if len(sourceImage) == 0 || len(targetImage) == 0 {
		return nil, errors.New(ErrEmptyBytes)
	}

	if similarity == 0.0 {
		return nil, errors.New(ErrBadSimilarity)
	}

	cmpFacesIn, cmpFacesInErr := rekognition.NewCompareFacesInput(
		sourceImage,
		targetImage,
		similarity,
	)
	if cmpFacesInErr != nil {
		return nil, cmpFacesInErr
	}

	valid := cmpFacesIn.Validate()
	if valid != nil {
		return nil, valid
	}

	cmpFacesOut, err := svc.RekognitionCompareFaces(cmpFacesIn)
	if err != nil {
		return nil, err
	}

	var out CompareFacesOutput

	unmarshalErr := UnmarshalCompareFacesOutput(*cmpFacesOut, &out)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &out, nil

}

// DetectFaces returns a *DetectFacesOutput
func (svc *Rekognition) DetectFaces(sourceImage []byte) (*DetectFacesOutput, error) {

	if len(sourceImage) == 0 {
		return nil, errors.New(ErrEmptyBytes)
	}

	dtcFacesIn, dtcFacesInErr := rekognition.NewDetectFacesInput(
		sourceImage,
	)
	if dtcFacesInErr != nil {
		return nil, dtcFacesInErr
	}

	valid := dtcFacesIn.Validate()
	if valid != nil {
		return nil, valid
	}

	dtcFacesOut, err := svc.RekognitionDetectFaces(dtcFacesIn)
	if err != nil {
		return nil, err
	}

	var out DetectFacesOutput

	unmarshalErr := UnmarshalDetectFacesOutput(*dtcFacesOut, &out)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &out, nil

}

// DetectText returns a *DetectTextOutput
func (svc *Rekognition) DetectText(sourceImage []byte) (*DetectTextOutput, error) {

	if len(sourceImage) == 0 {
		return nil, errors.New(ErrEmptyBytes)
	}

	dtcTextIn, dtcTextInErr := rekognition.NewDetectTextInput(
		sourceImage,
	)
	if dtcTextInErr != nil {
		return nil, dtcTextInErr
	}

	valid := dtcTextIn.Validate()
	if valid != nil {
		return nil, valid
	}

	dtcTextOut, err := svc.RekognitionDetectText(dtcTextIn)
	if err != nil {
		return nil, err
	}

	var out DetectTextOutput

	unmarshalErr := UnmarshalDetectTextOutput(*dtcTextOut, &out)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &out, nil

}

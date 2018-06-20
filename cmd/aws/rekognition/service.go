package rekognition

import "github.com/andream16/aws-sdk-go-bindings/pkg/aws/rekognition"

// CompareFaces returns a *CompareFacesOutput
func (svc *Rekognition) CompareFaces(sourceImage, targetImage []byte, similarity float64) (*CompareFacesOutput, error) {

	cmpFacesIn := rekognition.NewCompareFacesInput(
		sourceImage,
		targetImage,
		similarity,
	)

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

	dtcFacesIn := rekognition.NewDetectFacesInput(
		sourceImage,
	)

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

	dtcTextIn := rekognition.NewDetectTextInput(
		sourceImage,
	)

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

package rekognition

import "github.com/aws/aws-sdk-go/service/rekognition"

// CompareFacesInput embeds *rekognition.CompareFacesInput
type CompareFacesInput struct {
	*rekognition.CompareFacesInput
}

// CompareFacesOutput embeds *rekognition.CompareFacesOutput
type CompareFacesOutput struct {
	*rekognition.CompareFacesOutput
}

// DetectFacesInput embeds *rekognition.DetectFacesInput
type DetectFacesInput struct {
	*rekognition.DetectFacesInput
}

// DetectFacesOutput embeds *rekognition.DetectFacesOutput
type DetectFacesOutput struct {
	*rekognition.DetectFacesOutput
}

// DetectTextInput embeds *rekognition.DetectTextInput
type DetectTextInput struct {
	*rekognition.DetectTextInput
}

// DetectTextOutput embeds *rekognition.DetectTextOutput
type DetectTextOutput struct {
	*rekognition.DetectTextOutput
}

// RekognitionCompareFaces compares two faces returning their similarity
func (svc *Rekognition) RekognitionCompareFaces(input *CompareFacesInput) (*CompareFacesOutput, error) {

	compareFacesOut, err := svc.CompareFaces(input.CompareFacesInput)
	if err != nil {
		return nil, err
	}

	out := new(CompareFacesOutput)
	out.CompareFacesOutput = compareFacesOut

	return out, nil

}

// RekognitionDetectFaces detects faces in an image
func (svc *Rekognition) RekognitionDetectFaces(input *DetectFacesInput) (*DetectFacesOutput, error) {

	detectFacesOut, err := svc.DetectFaces(input.DetectFacesInput)
	if err != nil {
		return nil, err
	}

	out := new(DetectFacesOutput)
	out.DetectFacesOutput = detectFacesOut

	return out, nil

}

// RekognitionDetectText extracts text from an image
func (svc *Rekognition) RekognitionDetectText(input *DetectTextInput) (*DetectTextOutput, error) {

	detectTextOut, err := svc.DetectText(input.DetectTextInput)
	if err != nil {
		return nil, err
	}

	out := new(DetectTextOutput)
	out.DetectTextOutput = detectTextOut

	return out, nil

}

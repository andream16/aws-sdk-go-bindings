package rekognition

// RekognitionCompare compares two faces returning their similarity
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

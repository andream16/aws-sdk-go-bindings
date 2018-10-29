package rekognition

// RekognitionCompareFaces compares two faces returning their similarity
func (svc *Rekognition) RekognitionCompareFaces(sourceImage, targetImage []byte, similarity float64) (*CompareFacesOutput, error) {

	input, err := NewCompareFacesInput(
		sourceImage,
		targetImage,
		similarity,
	)
	if err != nil {
		return nil, err
	}

	valid := input.Validate()
	if valid != nil {
		return nil, valid
	}

	compareFacesOut, err := svc.CompareFaces(input)
	if err != nil {
		return nil, err
	}

	var out CompareFacesOutput

	err = UnmarshalCompareFacesOutput(compareFacesOut, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil

}

// RekognitionDetectFaces detects faces in an image
func (svc *Rekognition) RekognitionDetectFaces(sourceImage []byte) (*DetectFacesOutput, error) {

	input, err := NewDetectFacesInput(
		sourceImage,
	)
	if err != nil {
		return nil, err
	}

	valid := input.Validate()
	if valid != nil {
		return nil, valid
	}

	detectFacesOut, err := svc.DetectFaces(input)
	if err != nil {
		return nil, err
	}

	var out DetectFacesOutput

	err = UnmarshalDetectFacesOutput(detectFacesOut, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil

}

// RekognitionDetectText extracts text from an image
func (svc *Rekognition) RekognitionDetectText(sourceImage []byte) (*DetectTextOutput, error) {

	input, err := NewDetectTextInput(
		sourceImage,
	)
	if err != nil {
		return nil, err
	}

	valid := input.Validate()
	if valid != nil {
		return nil, valid
	}

	detectTextOut, err := svc.DetectText(input)
	if err != nil {
		return nil, err
	}

	var out DetectTextOutput

	err = UnmarshalDetectTextOutput(detectTextOut, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil

}

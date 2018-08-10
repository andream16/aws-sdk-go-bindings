package rekognition

// RekognitionCompareFaces compares two faces returning their similarity
func (svc *Rekognition) RekognitionCompareFaces(sourceImage, targetImage []byte, similarity float64) (*CompareFacesOutput, error) {

	input, inputErr := NewCompareFacesInput(
		sourceImage,
		targetImage,
		similarity,
	)
	if inputErr != nil {
		return nil, inputErr
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

	unmarshalErr := UnmarshalCompareFacesOutput(compareFacesOut, &out)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &out, nil

}

// RekognitionDetectFaces detects faces in an image
func (svc *Rekognition) RekognitionDetectFaces(sourceImage []byte) (*DetectFacesOutput, error) {

	input, inputErr := NewDetectFacesInput(
		sourceImage,
	)
	if inputErr != nil {
		return nil, inputErr
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

	unmarshalErr := UnmarshalDetectFacesOutput(detectFacesOut, &out)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &out, nil

}

// RekognitionDetectText extracts text from an image
func (svc *Rekognition) RekognitionDetectText(sourceImage []byte) (*DetectTextOutput, error) {

	input, inputErr := NewDetectTextInput(
		sourceImage,
	)
	if inputErr != nil {
		return nil, inputErr
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

	unmarshalErr := UnmarshalDetectTextOutput(detectTextOut, &out)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &out, nil

}

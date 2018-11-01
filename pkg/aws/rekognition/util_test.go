package rekognition

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/stretchr/testify/assert"
)

var compareFacesOutputMock = &rekognition.CompareFacesOutput{
	FaceMatches: []*rekognition.CompareFacesMatch{
		{
			Face: &rekognition.ComparedFace{
				BoundingBox: &rekognition.BoundingBox{
					Height: aws.Float64(0.43156373500823975),
					Left:   aws.Float64(0.37979263067245483),
					Top:    aws.Float64(0.09570241719484329),
					Width:  aws.Float64(0.288160115480423),
				},
				Confidence: aws.Float64(99.9919662475586),
				Landmarks: []*rekognition.Landmark{
					{
						Type: aws.String("eyeLeft"),
						X:    aws.Float64(0.47557365894317627),
						Y:    aws.Float64(0.28052255511283875),
					},
					{
						Type: aws.String("eyeRight"),
						X:    aws.Float64(0.5596158504486084),
						Y:    aws.Float64(0.2751060128211975),
					},
					{
						Type: aws.String("nose"),
						X:    aws.Float64(0.48886510729789734),
						Y:    aws.Float64(0.34084776043891907),
					},
					{
						Type: aws.String("mouthLeft"),
						X:    aws.Float64(0.4864692687988281),
						Y:    aws.Float64(0.4339553415775299),
					},
					{
						Type: aws.String("mouthRight"),
						X:    aws.Float64(0.5461938381195068),
						Y:    aws.Float64(0.4330655336380005),
					},
				},
				Pose: &rekognition.Pose{
					Pitch: aws.Float64(7.827611446380615),
					Roll:  aws.Float64(-0.8795318007469177),
					Yaw:   aws.Float64(-28.9353218078613287),
				},
				Quality: &rekognition.ImageQuality{
					Brightness: aws.Float64(43.768043518066406),
					Sharpness:  aws.Float64(99.9305191040039),
				},
			},
			Similarity: aws.Float64(98),
		},
	},
	SourceImageFace: &rekognition.ComparedSourceImageFace{
		BoundingBox: &rekognition.BoundingBox{
			Height: aws.Float64(0.3982542157173157),
			Left:   aws.Float64(0.29615384340286255),
			Top:    aws.Float64(0.1509365290403366),
			Width:  aws.Float64(0.2807692289352417),
		},
		Confidence: aws.Float64(99.97785949707031),
	},
	SourceImageOrientationCorrection: aws.String("ROTATE_0"),
	TargetImageOrientationCorrection: aws.String("ROTATE_0"),
	UnmatchedFaces:                   []*rekognition.ComparedFace{},
}

var detectFacesOutputMock = &rekognition.DetectFacesOutput{
	FaceDetails: []*rekognition.FaceDetail{
		{
			AgeRange: &rekognition.AgeRange{
				High: aws.Int64(60),
				Low:  aws.Int64(35),
			},
			Beard: &rekognition.Beard{
				Confidence: aws.Float64(99.0),
				Value:      aws.Bool(true),
			},
			Confidence: aws.Float64(99.0),
			Eyeglasses: &rekognition.Eyeglasses{
				Confidence: aws.Float64(99.0),
				Value:      aws.Bool(true),
			},
			EyesOpen: &rekognition.EyeOpen{
				Confidence: aws.Float64(99.0),
				Value:      aws.Bool(true),
			},
			Gender: &rekognition.Gender{
				Confidence: aws.Float64(99.0),
				Value:      aws.String("male"),
			},
		},
	},
}

var detectTextMock = &rekognition.DetectTextOutput{
	TextDetections: []*rekognition.TextDetection{
		{
			Confidence:   aws.Float64(99.0),
			Id:           aws.Int64(0),
			Type:         aws.String("WORD"),
			DetectedText: aws.String("Hello"),
		},
		{
			Confidence:   aws.Float64(99.0),
			Id:           aws.Int64(1),
			Type:         aws.String("WORD"),
			DetectedText: aws.String("World"),
		},
		{
			Confidence:   aws.Float64(99.0),
			Id:           aws.Int64(2),
			Type:         aws.String("WORD"),
			DetectedText: aws.String("!"),
		},
	},
}

func TestUnmarshalCompareFacesOutput(t *testing.T) {

	var out CompareFacesOutput

	err := UnmarshalCompareFacesOutput(compareFacesOutputMock, &out)

	assert.NoError(t, err)

	assert.NotEmpty(t, out)

}

func TestUnmarshalDetectFacesOutput(t *testing.T) {

	var out DetectFacesOutput

	err := UnmarshalDetectFacesOutput(detectFacesOutputMock, &out)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func TestUnmarshalDetectTextOutput(t *testing.T) {

	var out DetectTextOutput

	err := UnmarshalDetectTextOutput(detectTextMock, &out)

	assert.NoError(t, err)

	assert.NotEmpty(t, out)

}

func TestNewCompareFacesInput(t *testing.T) {

	target := []byte("someTarget")
	source := []byte("someSource")
	similarity := 90.0

	in, err := NewCompareFacesInput(source, target, similarity)

	assert.NoError(t, err)
	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.SourceImage.Bytes)
	assert.Equal(t, target, in.TargetImage.Bytes)
	assert.Equal(t, similarity, *in.SimilarityThreshold)

	_, err = NewCompareFacesInput([]byte{}, target, similarity)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

	_, err = NewCompareFacesInput(source, []byte{}, similarity)
	assert.Contains(t, err.Error(), ErrEmptyParameter)

	_, err = NewCompareFacesInput(source, target, 0.0)

	assert.Contains(t, err.Error(), ErrBadSimilarityParameter)

}

func TestNewDetectFacesInput(t *testing.T) {

	source := []byte("someSource")

	in, err := NewDetectFacesInput(source)

	assert.NoError(t, err)

	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.Image.Bytes)

}

func TestNewDetectTextInput(t *testing.T) {

	source := []byte("someSource")

	in, err := NewDetectTextInput(source)

	assert.NoError(t, err)

	assert.NotEmpty(t, in)
	assert.Equal(t, source, in.Image.Bytes)

}

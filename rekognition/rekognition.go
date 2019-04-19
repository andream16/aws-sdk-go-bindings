package rekognition

import (
	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
	"github.com/pkg/errors"
)

// Rekognitioner describes rekognition API.
type Rekognitioner interface {
	CompareFaces(
		source []byte,
		target []byte,
		treshold float64,
	) (*rekognition.CompareFacesOutput, error)
	DetectFaces(source []byte) (*rekognition.DetectFacesOutput, error)
	DetectText(source []byte) (*rekognition.DetectTextOutput, error)
}

// Rekognition is the alias for Rekognition.
type Rekognition struct {
	rekognition rekognitioniface.RekognitionAPI
}

// New returns a new S3.
func New(config *aws.Config) (*Rekognition, error) {

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize s3 session")
	}

	return &Rekognition{
		rekognition: rekognition.New(sess),
	}, nil

}

// CompareFaces compares two faces given a certain treshold.
func (r Rekognition) CompareFaces(
	source []byte,
	target []byte,
	treshold float64,
) (*rekognition.CompareFacesOutput, error) {

	if len(source) == 0 {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "source")
	}
	if len(target) == 0 {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "source")
	}

	in := &rekognition.CompareFacesInput{
		SourceImage: &rekognition.Image{
			Bytes: source,
		},
		TargetImage: &rekognition.Image{
			Bytes: target,
		},
		SimilarityThreshold: aws.Float64(treshold),
	}

	if err := in.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid input")
	}

	return r.rekognition.CompareFaces(in)

}

// DetectFaces detects faces in an image.
func (r Rekognition) DetectFaces(source []byte) (*rekognition.DetectFacesOutput, error) {

	if len(source) == 0 {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "source")
	}

	in := &rekognition.DetectFacesInput{
		Image: &rekognition.Image{
			Bytes: source,
		},
	}

	if err := in.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid input")
	}

	return r.rekognition.DetectFaces(in)

}

// DetectText detects text in an image.
func (r Rekognition) DetectText(source []byte) (*rekognition.DetectTextOutput, error) {

	if len(source) == 0 {
		return nil, errors.Wrap(bindings.ErrInvalidParameter, "source")
	}

	in := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			Bytes: source,
		},
	}

	if err := in.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid input")
	}

	return r.rekognition.DetectText(in)

}

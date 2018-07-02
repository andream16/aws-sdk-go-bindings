package rekognition

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/rekognition"
)

// Rekognition embeds *rekognition.Rekognition and is used to call rekognition methods on high level
type Rekognition struct {
	*rekognition.Rekognition
}

// New returns a *Rekognition given a *aws.Session. Region is optional.
func New(svc *aws.Session, region string) (*Rekognition, error) {

	rekSvc, err := rekognition.New(svc.Session, region)
	if err != nil {
		return nil, err
	}

	newRekSvc := new(Rekognition)
	newRekSvc.Rekognition = rekSvc

	return newRekSvc, nil

}

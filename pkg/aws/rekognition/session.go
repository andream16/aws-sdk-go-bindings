package rekognition

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
)

// Rekognition embeds *rekognition.Rekognition to be used to call New
type Rekognition struct {
	*rekognition.Rekognition
}

// New returns a new *Rekognition embedding *rekognition.Rekognition
func New(svc *pkgAws.Session, region string) (*Rekognition, error) {

	if len(region) > 0 {
		svc.Config.Region = aws.String(region)
	}

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	rekognitionSvc := new(Rekognition)
	rekognitionSvc.Rekognition = rekognition.New(newSvc)

	return rekognitionSvc, nil

}

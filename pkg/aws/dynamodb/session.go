package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
)

// DynamoDB embeds *dynamodb.DynamoDB
type DynamoDB struct {
	*dynamodb.DynamoDB
}

// New returns a new *DynamoDB
func New(svc *pkgAws.Session, endpoint string) (*DynamoDB, error) {

	if len(endpoint) > 0 {
		svc.Config.Endpoint = aws.String(endpoint)
	}

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	dynamoSvc := &DynamoDB{}
	dynamoSvc.DynamoDB = dynamodb.New(newSvc)

	return dynamoSvc, nil

}

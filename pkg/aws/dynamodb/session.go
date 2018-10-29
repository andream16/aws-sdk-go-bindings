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

	if endpoint != "" {
		svc.Config.Endpoint = aws.String(endpoint)
	}

	newSvc, err := session.NewSession(svc.Config)
	if err != nil {
		return nil, err
	}

	dynamoSvc := &DynamoDB{
		DynamoDB: dynamodb.New(newSvc),
	}

	return dynamoSvc, nil

}

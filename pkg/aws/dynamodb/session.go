package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB embeds *dynamodb.DynamoDB
type DynamoDB struct {
	*dynamodb.DynamoDB
}

func New(svc *aws.Session) (*DynamoDB, error) {

	newSvc, newSvcErr := session.NewSession(svc.Config)
	if newSvcErr != nil {
		return nil, newSvcErr
	}

	dynamoSvc := new(DynamoDB)
	dynamoSvc.DynamoDB = dynamodb.New(newSvc)

	return dynamoSvc, nil

}

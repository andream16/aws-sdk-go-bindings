package dynamodb

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/dynamodb"
)

// DynamoDB embeds *dynamodb.DynamoDB and is used to call sns methods on high level
type DynamoDB struct {
	*dynamodb.DynamoDB
}

// New returns a *DynamoDB given a *aws.Session
func New(svc *aws.Session) (*DynamoDB, error) {

	snsSvc, err := dynamodb.New(svc.Session)
	if err != nil {
		return nil, err
	}

	newSnsSvc := new(DynamoDB)
	newSnsSvc.DynamoDB = snsSvc

	return newSnsSvc, nil

}

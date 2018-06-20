package dynamodb

import "github.com/aws/aws-sdk-go/service/dynamodb"

// PutItemInput embeds *dynamodb.PutItemInput
type PutItemInput struct {
	*dynamodb.PutItemInput
}

// DynamoPutItem puts a given input on dynamodb
func (svc *DynamoDB) DynamoPutItem(input *PutItemInput) error {

	_, err := svc.PutItem(input.PutItemInput)
	if err != nil {
		return err
	}

	return nil

}

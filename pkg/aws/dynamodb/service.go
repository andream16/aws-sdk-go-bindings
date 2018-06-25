package dynamodb

import "github.com/aws/aws-sdk-go/service/dynamodb"

// PutItemInput embeds *dynamodb.PutItemInput
type PutItemInput struct {
	*dynamodb.PutItemInput
}

// GetItemInput embeds *dynamodb.GetItemInput
type GetItemInput struct {
	*dynamodb.GetItemInput
}

// GetItemOutput embeds *dynamodb.GetItemOutput
type GetItemOutput struct {
	*dynamodb.GetItemOutput
}

// DynamoPutItem puts a given input on dynamodb
func (svc *DynamoDB) DynamoPutItem(input *PutItemInput) error {

	_, err := svc.PutItem(input.PutItemInput)
	if err != nil {
		return err
	}

	return nil

}

// DynamoGetItem gets an item from DynamoDB given a valid *GetItemInput
func (svc *DynamoDB) DynamoGetItem(input *GetItemInput) (*GetItemOutput, error) {

	item, err := svc.GetItem(input.GetItemInput)
	if err != nil {
		return nil, err
	}

	out := new(GetItemOutput)
	out.GetItemOutput = item

	return out, nil

}

package dynamodb

import "github.com/andream16/aws-sdk-go-bindings/pkg/aws/dynamodb"

// PutItem puts a passed input into a passed table
func (svc *DynamoDB) PutItem(input interface{}, tableName string) error {

	in, err := dynamodb.NewPutItemInput(input, tableName)
	if err != nil {
		return err
	}

	putErr := svc.DynamoPutItem(in)
	if putErr != nil {
		return putErr
	}

	return nil

}

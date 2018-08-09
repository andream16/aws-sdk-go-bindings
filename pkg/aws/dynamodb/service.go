package dynamodb

// GetItemOutput embeds *dynamodb.GetItemOutput
type GetItemOutput interface{}

// DynamoPutItem puts a given input in a dynamodb table
func (svc *DynamoDB) DynamoPutItem(input interface{}, table string) error {

	newPutItemIn, newPutItemInErr := NewPutItemInput(input, table)
	if newPutItemInErr != nil {
		return newPutItemInErr
	}

	_, err := svc.PutItem(newPutItemIn)
	if err != nil {
		return err
	}

	return nil

}

// DynamoGetItem gets an item from DynamoDB given a key and its value.
// A *GetItemOutput will be returned
func (svc *DynamoDB) DynamoGetItem(table, keyName, keyValue string) (*GetItemOutput, error) {

	in, inErr := NewGetItemInput(
		table,
		keyName,
		keyValue,
	)
	if inErr != nil {
		return nil, inErr
	}

	item, err := svc.GetItem(in)
	if err != nil {
		return nil, err
	}

	out := new(GetItemOutput)

	itemErr := UnmarshalGetItemOutput(item, &out)
	if itemErr != nil {
		return nil, itemErr
	}

	return out, nil

}

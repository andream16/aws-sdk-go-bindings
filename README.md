[![Go Report Card](https://goreportcard.com/badge/github.com/AndreaM16/aws-sdk-go-bindings)](https://goreportcard.com/report/github.com/AndreaM16/aws-sdk-go-bindings) [![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/andream16/aws-sdk-go-bindings/blob/master/LICENSE.txt)
[![BCH compliance](https://bettercodehub.com/edge/badge/AndreaM16/aws-sdk-go-bindings?branch=master)](https://bettercodehub.com/)

# aws-sdk-go-bindings
Helper to easily access some [aws-sdk-go](https://github.com/aws/aws-sdk-go)'s methods. It also contains multiple methods to cover tricky problems like preparing an sns default message and unmarshal an image coming out from a stream like:

```
// UnmarshalStreamImage unmarshals a dynamo stream image in a pointer to an interface
func UnmarshalStreamImage(in map[string]events.DynamoDBAttributeValue, out interface{}) error {

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range in {

		bytes, err := v.MarshalJSON()
		if err != nil {
			return err
		}

		var dbAttr dynamodb.AttributeValue

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr

	}

	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)

}
```

At the moment it covers SNS, SQS, DynamoDB, Rekognition and S3.

## Utilization

You can simply import code from `pkg` package. Almost all the methods are exported so you can access them easily.

## Development

Install [dep](https://github.com/golang/dep) and run `dep ensure` inside the project's folder to get project's vendors.
If you want to fork it or just use it in local, edit `internal/configuration/configuration.json` as you wish. The default configuration contains endpoints to run the tests on [localstack](https://github.com/localstack/localstack). To run `Rekognition` tests you need to have an AWS account and use a region where the latter is available.

# aws-sdk-go-bindings [![CircleCI](https://circleci.com/gh/andream16/aws-sdk-go-bindings.svg?style=svg)](https://circleci.com/gh/andream16/aws-sdk-go-bindings) [![GoDoc](https://godoc.org/github.com/andream16/aws-sdk-go-bindings?status.svg)](https://godoc.org/github.com/andream16/aws-sdk-go-bindings) [![Go Report Card](https://goreportcard.com/badge/github.com/AndreaM16/aws-sdk-go-bindings)](https://goreportcard.com/report/github.com/AndreaM16/aws-sdk-go-bindings) [![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/andream16/aws-sdk-go-bindings/blob/master/LICENSE.txt)

Helper to easily access some [aws-sdk-go](https://github.com/aws/aws-sdk-go)'s methods and lambda utilities like preparing an sns default message and unmarshal an image coming out from a stream like:

```
// UnmarshalDynamoEvent unmarshals a events.DynamoDBEventRecord into a given target.
// Out has to be a pointer.
func UnmarshalDynamoEvent(event events.DynamoDBEventRecord, out interface{}) error {

	if reflect.DeepEqual(event, reflect.Zero(reflect.TypeOf(event)).Interface()) {
		return errors.Wrap(bindings.ErrInvalidParameter, "event")
	}
	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		return errors.Wrap(bindings.ErrInvalidParameter, "out")
	}

	img := event.Change.NewImage
	if len(img) == 0 {
		return errors.New("event's image is empty")
	}

	m := make(map[string]*dynamodb.AttributeValue, len(img))

	for k, v := range img {

		b, err := v.MarshalJSON()
		if err != nil {
			return errors.Wrap(err, "unable to marshal current element to json")
		}

		var attr dynamodb.AttributeValue

		err = json.Unmarshal(b, &attr)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal current element to json")
		}

		m[k] = &attr

	}

	return dynamodbattribute.UnmarshalMap(m, out)

}
```

## Contributors

Special thanks to:

 - [martingallagher](https://github.com/martingallagher/)
 - [ferruvich](https://github.com/ferruvich/)

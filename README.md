[![Go Report Card](https://goreportcard.com/badge/github.com/AndreaM16/aws-sdk-go-bindings)](https://goreportcard.com/report/github.com/AndreaM16/aws-sdk-go-bindings) [![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/andream16/aws-sdk-go-bindings/blob/master/LICENSE.txt)

# aws-sdk-go-bindings
Helper to easily access some [aws-sdk-go](https://github.com/aws/aws-sdk-go)'s methods. It also contains multiple methods to cover tricky problems like preparing and sns default message and unmarshal an image coming out from a stream. At the moment it covers SNS, DynamoDB, Rekognition and S3.

## Utilization

You can import code from both `cmd` and `pkg` root package. `cmd` contains high level methods that interact with `pkg` ones while `pkg` directly interacts with aws-sdk-go.

## Development

Install [dep](https://github.com/golang/dep) and run `dep ensure` inside the project's folder to get project's vendors.
If you want to fork it or just use it in local, edit `internal/configuration/configuration.json` by setting your aws options like:

```
{
  "region" : "eu-central-1",
  "SNS" : {
    "target_arn" : "arn:aws:sns:eu-central-1:${your_aws_account_id}:${your_resource_name}"
  },
  "DynamoDB" : {
    "endpoint" : "your_local_dynamo_endpoint",
    "pkg_table_name" : "some_table_1",
    "cmd_table_name" : "some_table_2",
    "primary_key" : "some_param"
  },
  "S3" : {
      "bucket" : "your_bucket",
      "source_image" : "path_to_a_test_image"
    }
}
```

Also, to make `dynamodb` tests work you need to start a [local dynamodb](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html)
and run it on the same endpoint you put in the configuration above like `java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb -port 4200
` if you used `"DynamoDB.endpoint" : "http://localhost:4200/"`.

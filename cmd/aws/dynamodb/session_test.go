package dynamodb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc, svcErr := aws.New(cfg.Region)

	assert.NoError(t, svcErr)

	dynamoSvc, dynamoSvcErr := New(svc, cfg.DynamoDB.Endpoint)

	assert.NoError(t, dynamoSvcErr)
	assert.NotEmpty(t, dynamoSvc)

}

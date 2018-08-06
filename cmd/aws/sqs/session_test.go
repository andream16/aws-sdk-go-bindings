package sqs

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

	snsSvc, snsSvcErr := New(svc)

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

}

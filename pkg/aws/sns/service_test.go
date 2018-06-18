package sns

import (
	"testing"
	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/aws"
)

func TestSnsPublish(t *testing.T) {

	awsSvc, awsSvcErr := pkgAws.New()

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	snsSvc, snsSvcErr := New(awsSvc.Session)

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

	in := &PublishInput{
		PublishInput: &sns.PublishInput{
			Message:          aws.String(`{"default":"{\"par1\":\"pr1\",\"par2\":\"pr2\"}"}`),
			TargetArn:        aws.String(`some_endpoint`),
			MessageStructure: aws.String(`json`),
		},
	}

	err := snsSvc.SnsPublish(in)

	assert.NoError(t, err)

}


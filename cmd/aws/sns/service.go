package sns

import (
	"errors"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/sns"
	"reflect"
)

// Publish publishes a given input to in a given targetArn
func (svc *SNS) Publish(input interface{}, targetArn string) error {

	if reflect.DeepEqual(input, reflect.Zero(reflect.TypeOf(input)).Interface()) || targetArn == "" {
		return errors.New(ErrEmptyParameter)
	}

	out, outErr := sns.NewPublishInput(
		input,
		targetArn,
	)
	if outErr != nil {
		return outErr
	}

	pubErr := svc.SnsPublish(out)
	if pubErr != nil {
		return pubErr
	}

	return nil

}

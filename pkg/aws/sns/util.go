package sns

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/service/sns"
)

const messageStructure = "json"

// Body is used to initialize a valid SNS message
type Body struct {
	Default string `json:"default"`
}

// NewPublishInput returns a new *PublishInput given a body and an endpoint
func NewPublishInput(body interface{}, endpoint string) (*PublishInput, error) {

	if endpoint == "" {
		return nil, errors.New(ErrEmptyParameter)
	}

	inBytes, inErr := json.Marshal(body)
	if inErr != nil {
		return nil, inErr
	}

	// Mandatory since SNS needs escaped `"`. So we need to escape them to `\"`
	unquote := strings.Replace(string(inBytes), `"`, "\"", -1)

	// Mandatory since SNS needs bodies like:
	// {
	// 		"default" : {
	// 			\"par1\" : \"some value\"
	// 		}
	// }
	snsBody := Body{
		Default: unquote,
	}

	// Mandatory since we want to get a string out of encoded bytes
	msgBytes, msgErr := json.Marshal(snsBody)
	if msgErr != nil {
		return nil, msgErr
	}

	publishInput := new(sns.PublishInput)
	publishInput = publishInput.SetMessage(string(msgBytes))
	publishInput = publishInput.SetMessageStructure(messageStructure)
	publishInput = publishInput.SetTargetArn(endpoint)

	out := new(PublishInput)
	out.PublishInput = publishInput

	return out, nil

}

package sns

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/sns"

	intErr "github.com/andream16/aws-sdk-go-bindings/internal/error"
)

// Body is used to initialize a valid SNS message
type Body struct {
	Default string `json:"default"`
}

// NewPublishInput returns a new *PublishInput given a body and an endpoint
func NewPublishInput(input interface{}, endpoint string) (*sns.PublishInput, error) {

	if len(endpoint) == 0 {
		return nil, intErr.Format(Endpoint, ErrEmptyParameter)
	}

	if reflect.ValueOf(input).Kind() == reflect.Ptr {
		return nil, intErr.Format(Input, ErrPointerParameterNotAllowed)
	}

	inBytes, inErr := json.Marshal(input)
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

	out := &sns.PublishInput{}
	out = out.SetMessage(string(msgBytes))
	out = out.SetMessageStructure(MessageStructure)
	out = out.SetTargetArn(endpoint)

	return out, nil

}

// UnmarshalMessage unmarshal an SNS Message to a given interface
func UnmarshalMessage(message string, input interface{}) error {

	if len(message) == 0 {
		return intErr.Format(Message, ErrEmptyParameter)
	}

	if reflect.ValueOf(input).Kind() != reflect.Ptr {
		return intErr.Format(Input, ErrNoPointerParameter)
	}

	uS := unescapeMessageString(message)

	unmarshalErr := json.Unmarshal([]byte(uS), input)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil

}

// unescapeMessageString takes a SNS message string like
// `"{\"stuff\" : \"somevalue\"}"` and outputs `"{"stuff" : "somevalue"}"`
func unescapeMessageString(in string) string {
	return strings.Replace(in, `\"`, `"`, -1)
}

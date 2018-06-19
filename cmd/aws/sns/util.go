package sns

import (
	"encoding/json"
	"strings"
)

// UnmarshalMessage unmarshal an SNS Message to a given interface
func UnmarshalMessage(msg string, in interface{}) error {

	uS := unescapeMessageString(msg)

	unmarshalErr := json.Unmarshal([]byte(uS), in)
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

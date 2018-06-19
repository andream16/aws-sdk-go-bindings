package sns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestSNSMessage struct {
	Param1 string `json:"param1"`
}

func TestMessageToJson(t *testing.T) {

	var p TestSNSMessage

	str := `{\"param1\":\"some_value\"}`

	err := UnmarshalMessage(str, &p)

	assert.NoError(t, err)
	assert.Equal(t, "some_value", p.Param1)

}

func TestUnescapeMessageString(t *testing.T) {

	str := `{\"some_param\":\"some_value\"}`

	s := unescapeMessageString(str)

	assert.Equal(t, `{"some_param":"some_value"}`, s)

}

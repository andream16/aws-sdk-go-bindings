package sns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestSNSMessage struct {
	Param1 string `json:"param1"`
}

func TestUnmarshalMessage(t *testing.T) {

	var p TestSNSMessage

	str := `{\"param1\":\"some_value\"}`

	err := UnmarshalMessage(str, &p)

	assert.NoError(t, err)
	assert.Equal(t, "some_value", p.Param1)

	shouldBeErr1 := UnmarshalMessage("", &p)

	assert.Error(t, shouldBeErr1)
	assert.Equal(t, ErrEmptyParameter, shouldBeErr1.Error())

	shouldBeErr2 := UnmarshalMessage("some_val", p)

	assert.Error(t, shouldBeErr2)
	assert.Equal(t, ErrNoPointerParameter, shouldBeErr2.Error())

}

func TestUnescapeMessageString(t *testing.T) {

	str := `{\"some_param\":\"some_value\"}`

	s := unescapeMessageString(str)

	assert.Equal(t, `{"some_param":"some_value"}`, s)

}

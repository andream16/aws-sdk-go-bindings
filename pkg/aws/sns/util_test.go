package sns

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestBuildPublishInputType struct {
	Par1 string `json:"par1"`
	Par2 string `json:"par2"`
}

type TestSNSMessage struct {
	Param1 string `json:"param1"`
}

func TestBuildPublishInput(t *testing.T) {

	body := []byte(`
	{
		"par1" : "pr1",
		"par2" : "pr2"
	}
	`)

	var testB TestBuildPublishInputType

	e := json.Unmarshal(body, &testB)

	assert.NoError(t, e)

	res, err := NewPublishInput(testB, "edp")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	assert.Equal(t, "edp", *res.TargetArn)
	assert.NotEqual(t, 0, len(*res.Message))

	_, errEmptyParameter := NewPublishInput(testB, "")

	assert.Error(t, errEmptyParameter)
	assert.Contains(t, errEmptyParameter.Error(), ErrEmptyParameter)

}

func TestUnmarshalMessage(t *testing.T) {

	var p TestSNSMessage

	str := `{\"param1\":\"some_value\"}`

	err := UnmarshalMessage(str, &p)

	assert.NoError(t, err)
	assert.Equal(t, "some_value", p.Param1)

	shouldBeErr1 := UnmarshalMessage("", &p)

	assert.Error(t, shouldBeErr1)
	assert.Contains(t, shouldBeErr1.Error(), ErrEmptyParameter)

	shouldBeErr2 := UnmarshalMessage("some_val", p)

	assert.Error(t, shouldBeErr2)
	assert.Contains(t, shouldBeErr2.Error(), ErrNoPointerParameter)

}

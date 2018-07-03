package sns

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestBuildPublishInputType struct {
	Par1 string `json:"par1"`
	Par2 string `json:"par2"`
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

	assert.Equal(t, "edp", *res.PublishInput.TargetArn)
	assert.NotEqual(t, 0, len(*res.PublishInput.Message))

	_, errEmptyParameter := NewPublishInput(testB, "")

	assert.Error(t, errEmptyParameter)
	assert.Equal(t, ErrEmptyParameter, errEmptyParameter.Error())

}

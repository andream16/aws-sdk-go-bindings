package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {

	err := "SomeError"
	val := "SomeParameterName"

	s := Format(val, err)

	assert.Contains(t, s.Error(), err, val)

}

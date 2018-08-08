package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatError(t *testing.T) {

	err := "SomeError"
	val := "SomeParameterName"

	s := FormatError(val, err)

	assert.Contains(t, s.Error(), err, val)

}

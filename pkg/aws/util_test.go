package aws

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewSessionInput(t *testing.T) {

	region := "some_region"

	out := NewSessionInput(region)

	assert.Equal(t, region, out.region)

}

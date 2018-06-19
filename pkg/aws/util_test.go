package aws

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSessionInput(t *testing.T) {

	region := "some_region"

	out := NewSessionInput(region)

	assert.Equal(t, region, out.region)

}

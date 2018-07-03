package aws

import (
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	in := NewSessionInput(cfg.Region)

	assert.NotEmpty(t, in)

	svc, err := New(in)

	assert.NotEmpty(t, svc)
	assert.NoError(t, err)

	_, errNoRegionProvided := New(&SessionInput{
		region: "",
	})

	assert.Error(t, errNoRegionProvided)
	assert.Equal(t, ErrNoRegionProvided, errNoRegionProvided.Error())

}

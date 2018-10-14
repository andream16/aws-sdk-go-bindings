package aws

import "errors"

// SessionInput contains the input to be passed to New
type SessionInput struct {
	region string
}

// NewSessionInput returns a new *SessionInput
func NewSessionInput(region string) (*SessionInput, error) {

	if region == "" {
		return nil, errors.New(ErrNoRegionProvided)
	}

	svc := &SessionInput{
		region: region,
	}

	return svc, nil

}

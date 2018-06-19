package aws

// SessionInput contains the input to be passed to New
type SessionInput struct {
	region string
}

// NewSessionInput returns a new *SessionInput
func NewSessionInput(region string) *SessionInput {

	svc := new(SessionInput)

	svc.region = region

	return svc

}

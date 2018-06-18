package aws

type SessionInput struct {
	region string
}

// NewSessionInput returns a new *SessionInput
func NewSessionInput(region string) *SessionInput {

	svc := new(SessionInput)

	svc.region = region

	return svc

}

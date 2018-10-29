package sns

// SnsPublish publishes an input on a given SNS targetArn
func (svc *SNS) SnsPublish(input interface{}, targetArn string) (err error) {

	in, err := NewPublishInput(
		input,
		targetArn,
	)
	if err != nil {
		return err
	}

	_, err = svc.SNS.Publish(in)
	if err != nil {
		return err
	}

	return nil

}

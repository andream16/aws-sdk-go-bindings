package s3

// S3CreateBucket creates a new bucket given a bucketName
func (svc *S3) S3CreateBucket(bucketName string) error {

	in, inErr := NewCreateBucketInput(bucketName)
	if inErr != nil {
		return inErr
	}

	_, err := svc.S3.CreateBucket(in)
	if err != nil {
		return err
	}

	return nil
}

// S3GetObject retrieves an object from S3 given a bucket name and a source image
func (svc *S3) S3GetObject(bucketName, sourceImage string) ([]byte, error) {

	s3In, s3InErr := NewGetObjectInput(
		bucketName,
		sourceImage,
	)
	if s3InErr != nil {
		return nil, s3InErr
	}

	getObjectOut, err := svc.GetObject(s3In)
	if err != nil {
		return nil, err
	}

	out, err := UnmarshalGetObjectOutput(getObjectOut)
	if err != nil {
		return nil, err
	}

	return out, nil

}

// S3PutObject puts a given object on S3
func (svc *S3) S3PutObject(bucketName, objectName, objectPath string) error {

	imgMeta, readErr := ReadImage(objectPath)
	if readErr != nil {
		return readErr
	}

	in, inErr := NewPutObjectInput(
		bucketName,
		objectName,
		imgMeta.ContentType,
		imgMeta.Body,
		imgMeta.ContentSize,
	)
	if inErr != nil {
		return inErr
	}

	_, err := svc.S3.PutObject(in)
	if err != nil {
		return err
	}

	return nil

}

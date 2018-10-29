package main

import (
	"fmt"
	"log"

	"github.com/andream16/aws-sdk-go-bindings/internal/configuration"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
)

const (

	// S3OBJECTPATH path to an example image
	S3OBJECTPATH = "assets/compare_faces_test-source.jpg"
)

func main() {

	// Getting Configuration
	cfg, err := configuration.Get()
	if err != nil {
		log.Fatal(err)
	}

	// Initializing a new AWS Session
	svcIn, err := aws.NewSessionInput(cfg.Region)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := aws.New(svcIn)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing a new S3 Session (Endpoint is optional, can use also `""`)
	s3Svc, err := s3.New(svc, cfg.S3.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Bucket, don't catch error so that we can try creating it over and over
	createBucketIn, err := s3.NewCreateBucketInput(cfg.S3.Bucket)
	if err != nil {
		log.Fatal(err)
	}

	s3Svc.CreateBucket(createBucketIn)

	fmt.Println(fmt.Sprintf("successfully created bucket %v", cfg.S3.Bucket))

	// Put an object
	if err := s3Svc.S3PutObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
		S3OBJECTPATH,
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("successfully put object %v in bucket %v", S3OBJECTPATH, cfg.S3.Bucket))

	// Get Object
	_, err = s3Svc.S3GetObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("successfully got object %v", cfg.S3.SourceImage))

}

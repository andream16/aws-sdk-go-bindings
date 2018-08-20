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
	cfg, cfgErr := configuration.Get()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	// Initializing a new AWS Session
	svcIn, svcInErr := aws.NewSessionInput(cfg.Region)
	if svcInErr != nil {
		log.Fatal(svcInErr)
	}

	svc, svcErr := aws.New(svcIn)
	if svcErr != nil {
		log.Fatal(svcErr)
	}

	// Initializing a new S3 Session (Endpoint is optional, can use also `""`)
	s3Svc, s3SvcErr := s3.New(svc, cfg.S3.Endpoint)
	if s3SvcErr != nil {
		log.Fatal(s3SvcErr)
	}

	// Create a Bucket, don't catch error so that we can try creating it over and over
	createBucketIn, createBucketInErr := s3.NewCreateBucketInput(cfg.S3.Bucket)
	if createBucketInErr != nil {
		log.Fatal(createBucketInErr)
	}

	s3Svc.CreateBucket(createBucketIn)

	fmt.Println(fmt.Sprintf("successfully created bucket %v", cfg.S3.Bucket))

	// Put an object
	if putObjectErr := s3Svc.S3PutObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
		S3OBJECTPATH,
	); putObjectErr != nil {
		log.Fatal(putObjectErr)
	}

	fmt.Println(fmt.Sprintf("successfully put object %v in bucket %v", S3OBJECTPATH, cfg.S3.Bucket))

	// Get Object
	_, objectErr := s3Svc.S3GetObject(
		cfg.S3.Bucket,
		cfg.S3.SourceImage,
	)
	if objectErr != nil {
		log.Fatal(objectErr)
	}

	fmt.Println(fmt.Sprintf("successfully got object %v", cfg.S3.SourceImage))

}

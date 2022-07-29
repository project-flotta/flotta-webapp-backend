package storage

import (
	"context"
	"fmt"
	"github.com/ahmadateya/flotta-webapp-backend/config"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

type S3 struct {
	awsClient *s3.Client
}

// configS3 creates the S3 client
func initS3Client() S3 {
	// read configurations from env file
	cfg, _ := config.NewConfig("./config.yaml")

	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(cfg.Storage.S3.Region),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.Storage.S3.AccessKeyId,
				cfg.Storage.S3.SecretAccessKey,
				"")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(awsCfg)
	return S3{
		awsClient: client,
	}
}

// ListS3TopLevelFolders returns the list of top level folders in the S3 bucket
// Top Level folders is supposed to represent the machines registered in Flotta
func (s *S3) ListS3TopLevelFolders() {
	cfg, _ := config.NewConfig("./config.yaml")

	input := &s3.ListObjectsV2Input{
		Bucket: &cfg.Storage.S3.Bucket,
	}

	resp, err := s.awsClient.ListObjectsV2(context.TODO(), input)

	if err != nil {
		fmt.Printf("Got error retrieving list of objects: %v\n", err)
		return
	}

	fmt.Printf("############### resp %v\n", resp)

	//for _, item := range resp.Contents {
	//	fmt.Println("Name:          ", *item.Key)
	//	fmt.Println("Last modified: ", *item.LastModified)
	//	fmt.Println("Size:          ", item.Size)
	//	fmt.Println("Storage class: ", item.StorageClass)
	//	fmt.Println("")
	//}
	//
	//fmt.Println("Found", len(resp.Contents), "items in bucket", *bucket)
	//fmt.Println("")
}

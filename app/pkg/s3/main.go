package s3

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

// InitS3Client initializes the S3 client
func InitS3Client() S3 {
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

// ListTopLevelFolders returns the list of top level folders in the S3 bucket
// Top Level folders is supposed to represent the machines registered in Flotta
func (s *S3) ListTopLevelFolders() []string {
	cfg, err := config.NewConfig("./config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file : %v", err)
	}

	delimiter := "/"
	resp, err := s.awsClient.ListObjectsV2(
		context.TODO(),
		&s3.ListObjectsV2Input{
			Bucket:    &cfg.Storage.S3.Bucket,
			Delimiter: &delimiter,
		})

	if err != nil {
		fmt.Printf("Got error retrieving list of objects: %v\n", err)
	}

	topLevelFolders := make([]string, 0)
	for _, obj := range resp.CommonPrefixes {
		topLevelFolders = append(topLevelFolders, *obj.Prefix)
	}

	return topLevelFolders
}

package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gophero/goal/assert"
	"io"
	"log"
)

type Conf struct {
	Region       string
	AccessKey    string
	AccessSecret string
}

type Client struct {
	*s3.Client
}

func NewS3(conf *Conf) *Client {
	return conn(conf.Region, conf.AccessKey, conf.AccessSecret)
}

func conn(region, accessKey, accessSecret string) *Client {
	cre := &credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: accessSecret,
		},
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithCredentialsProvider(cre))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &Client{Client: s3.NewFromConfig(cfg)}
}

func (client *Client) UploadFile(bucketName string, objectKey string, file io.Reader) error {
	assert.Require(bucketName != "", "bucket is required")
	assert.Require(file != nil, "file is required")

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("put object %v to %v failed, err: %w", objectKey, bucketName, err)
	}
	return nil
}

func (client *Client) DeleteFile(bucketName string, objectKey string) error {
	assert.Require(bucketName != "", "bucket is required")
	assert.Require(objectKey != "", "objectKey is required")
	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &bucketName, Key: &objectKey,
	})
	if err != nil {
		return fmt.Errorf("dekete object %v failed from %v, error: %w", objectKey, bucketName, err)
	}
	return nil
}

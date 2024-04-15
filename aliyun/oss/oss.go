package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gophero/goal/assert"
	"github.com/gophero/goal/errorx"
	"io"
)

type Conf struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

type Client struct {
	*oss.Client
}

func NewOss(conf *Conf) *Client {
	return &Client{Client: conn(conf.Endpoint, conf.AccessKey, conf.AccessSecret)}
}

func conn(endpoint, accessKey, accessSecret string, options ...oss.ClientOption) *oss.Client {
	fmt.Printf("Aliyun OSS Go SDK Version: %s\n", oss.Version)
	client, err := oss.New(endpoint, accessKey, accessSecret, options...)
	if err != nil {
		panic(err)
	}
	return client
}

func (client *Client) UploadFile(bucketName string, objectKey string, file io.Reader) error {
	assert.Require(client != nil, "oss.Client is required")
	assert.Require(bucketName != "", "bucket is required")
	assert.Require(file != nil, "file is required")

	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 上传文件
	return bucket.PutObject(objectKey, file)
}

func (client *Client) DeleteFile(bucketName string, objectKey string) error {
	assert.Require(client != nil, "oss.Client is required")
	assert.Require(bucketName != "", "bucket is required")
	assert.Require(objectKey != "", "objectKey is required")

	bucket, err := client.Bucket(bucketName)
	errorx.Throw(err)
	return bucket.DeleteObject(objectKey)
}

package api

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	id     int64
	bucket string
	region string
	conf   string
	svc    *s3.S3
}

// NewClient create new S3Client
func NewS3Client(conf string) (ClientAPI, error) {
	awsConf := aws.NewConfig().
		WithRegion(conf).
		WithMaxRetries(3)

	sessionObj, err := session.NewSession(awsConf)
	if err != nil {
		return nil, errors.New("cannot create S3 client")
	}
	return &S3Client{svc: s3.New(sessionObj)}, nil
}

// GetObject download key
func (sc *S3Client) GetObject(key string) ([]byte, error) {
	out, err := sc.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(sc.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(out.Body)
	defer func() { _ = out.Body.Close() }()
	if err != nil {
		return nil, err
	}
	if len(data) < int(*out.ContentLength) {
		return nil, fmt.Errorf("aws s3 short read for %s: expect %d bytes, got %d bytes", key, *out.ContentLength, len(data))
	}

	return data, nil
}

// PutObject upload object
func (sc *S3Client) PutObject(key string, payload []byte) error {
	_, err := sc.svc.PutObject(&s3.PutObjectInput{
		ACL:    aws.String(s3.ObjectCannedACLBucketOwnerFullControl),
		Bucket: aws.String(sc.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(payload),
	})
	return err
}

// DeleteObject delete key
func (sc *S3Client) DeleteObject(key string) error {
	_, err := sc.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(sc.bucket),
		Key:    aws.String(key),
	})
	return err
}

// ListObjects list object key
func (sc *S3Client) ListObjects(prefix string) (interface{}, error) {
	result, err := sc.svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(sc.bucket),
		Prefix: aws.String(prefix),
	})
	return result, err
}

// ExistObject make use of HeadObject to check existing of an object
func (sc *S3Client) ExistObject(key string) bool {
	_, err := sc.svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(sc.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return false
	}
	return true
}

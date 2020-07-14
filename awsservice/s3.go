package awsservice

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 Service Client Operator
type S3 struct {
	Client *s3.S3
}

// NewS3 Create S3 Client
func NewS3() S3 {
	// Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create Client from SDK
	client := s3.New(sess)
	return S3{
		Client: client,
	}
}

// ListBuckets S3 ListOperation Wrapper
func (s S3) ListBuckets() (*s3.ListBucketsOutput, error) {
	listBuckets, err := s.Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return listBuckets, nil
}

// ListObjects S3 ListObjects Wrapper
func (s S3) ListObjects(bucket string) (*s3.ListObjectsOutput, error) {
	listObjects, err := s.Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		return nil, err
	}
	return listObjects, nil
}

// PutObject S3 PutObject Wrapper
func (s S3) PutObject(filePath string, bucket, string, objectPath string) (*s3.PutObjectOutput, error) {
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader(filePath)),
		Bucket: aws.String(bucket),
		Key:    aws.String(objectPath),
	}
	result, err := s.Client.PutObject(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

package awsservice

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rivo/tview"
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

// SetObjectsLists Set Objects To Tview's List
func (s S3) SetObjectList(objectList *tview.List, bucketsList *tview.List, app *tview.Application) error {
	objectList.Clear()
	currentItem := bucketsList.GetCurrentItem()
	itemText, _ := bucketsList.GetItemText(currentItem)
	listObjects, err := s.Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(itemText)})
	if err != nil {
		return err
	}
	for _, object := range listObjects.Contents {
		objectList.AddItem(*object.Key, object.LastModified.Format("2006/01/02 15:04:05"), 0, nil)
	}
	app.SetFocus(objectList)
	return nil
}

// SetBucketsList Set Buckets To Tview's List
func (s S3) SetBucketsList(bucketsList *tview.List, serviceList *tview.List, objectList *tview.List, app *tview.Application) error {
	go objectList.Clear()
	bucketsList.Clear()
	listBuckets, err := s.Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}
	for _, bucket := range listBuckets.Buckets {
		bucketsList.AddItem(*bucket.Name, *listBuckets.Owner.DisplayName, 0, func() {
			s.SetObjectList(objectList, bucketsList, app)
		})
	}
	app.SetFocus(bucketsList)
	return nil
}

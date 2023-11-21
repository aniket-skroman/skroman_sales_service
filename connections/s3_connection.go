package connections

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Connection interface {
	MakeNewSession() (*session.Session, error)
	GetBucketName() string
	UploadOrderQuatation(file multipart.File, handler *multipart.FileHeader) (string, error)
	DeleteFiles(file_path string) error
}

type s3_connection struct {
	AccessKey  string
	SecretKey  string
	Region     string
	BucketName string
}

func NewS3Connection() S3Connection {
	return &s3_connection{
		AccessKey:  "AKIA3VMV3LWIR6TTGJGK",
		SecretKey:  "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",
		Region:     "ap-south-1",
		BucketName: "skromansupportbucket",
	}
}

func (s3_bucket *s3_connection) MakeNewSession() (*session.Session, error) {
	creds := credentials.NewStaticCredentials(s3_bucket.AccessKey,
		s3_bucket.SecretKey, "",
	)

	_, err := creds.Get()

	if err != nil {
		return nil, err
	}

	cfg := aws.NewConfig().WithRegion(s3_bucket.Region).WithCredentials(creds)
	sess, _ := session.NewSession(cfg)

	return sess, nil
}

func (s3_bucket *s3_connection) GetBucketName() string {
	return s3_bucket.BucketName
}

func (s3_bucket *s3_connection) UploadOrderQuatation(file multipart.File, handler *multipart.FileHeader) (string, error) {
	defer file.Close()

	file_name := handler.Filename

	sess, err := s3_bucket.MakeNewSession()

	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(sess)

	file_name = "quatation/" + file_name

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3_bucket.GetBucketName()), // Bucket to be used
		Key:    aws.String(file_name),                 // Name of the file to be saved
		Body:   file,                                  // File
	})

	return file_name, err
}

func (s3_bucket *s3_connection) DeleteFiles(file_path string) error {
	sess, err := s3_bucket.MakeNewSession()

	if err != nil {
		return err
	}

	svc := s3.New(sess)

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s3_bucket.GetBucketName()),
		Key:    aws.String(file_path),
	})

	return err
}

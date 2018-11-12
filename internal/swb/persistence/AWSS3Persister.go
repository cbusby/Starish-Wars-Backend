package persistence

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// AWSS3Persister implementation of Persister that reads/writes to Amazon S3
type AWSS3Persister struct {
}

// Save Save content to Amazon S3
func (a AWSS3Persister) Save(name string, contents string) error {
	region := aws.String("us-east-2")
	bucket := aws.String("aluminum-falcon")

	sess, _ := session.NewSession(&aws.Config{Region: region})
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: bucket,
		Key:    aws.String("games/" + name + ".json"),
		Body:   strings.NewReader(contents),
	})
	return err
}

func (a AWSS3Persister) Read(name string) (string, error) {
	region := aws.String("us-east-2")
	bucket := aws.String("aluminum-falcon")
	sess, _ := session.NewSession(&aws.Config{Region: region})
	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: bucket,
		Key:    aws.String("games/" + name + ".json"),
	}
	result, err := svc.GetObject(input)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", errors.New("not found")
	}
	buf := bytes.NewBuffer(nil)
	n, copyErr := io.Copy(buf, result.Body)
	if copyErr != nil {
		return "", copyErr
	}
	contents := string(buf.Bytes()[:n])
	return contents, nil
}

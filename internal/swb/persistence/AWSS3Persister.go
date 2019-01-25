package persistence

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// AWSS3Persister implementation of Persister that reads/writes to Amazon S3
type AWSS3Persister struct {
}

var region = aws.String(os.Getenv("REGION"))
var bucket = aws.String(os.Getenv("BUCKET"))
var filenameTemplate = "games/%s.json"

// Save Save content to Amazon S3
func (a AWSS3Persister) Save(gameID string, contents string) error {
	sess, _ := getAWSSession()
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: bucket,
		Key:    getSWBFilename(gameID),
		Body:   strings.NewReader(contents),
	})
	return err
}

func (a AWSS3Persister) Read(gameID string) (string, error) {
	sess, _ := getAWSSession()
	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: bucket,
		Key:    getSWBFilename(gameID),
	}
	result, err := svc.GetObject(input)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", fmt.Errorf("Not found: %s", gameID)
	}
	buf := bytes.NewBuffer(nil)
	n, copyErr := io.Copy(buf, result.Body)
	if copyErr != nil {
		return "", copyErr
	}
	contents := string(buf.Bytes()[:n])
	return contents, nil
}

func getAWSSession() (*session.Session, error) {
	return session.NewSession(&aws.Config{Region: region})
}

func getSWBFilename(gameID string) *string {
	return aws.String(fmt.Sprintf(filenameTemplate, gameID))
}

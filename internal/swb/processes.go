package swb

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// Create creates a new game, writes the state to S3, and returns the GameID and the state
func Create() (string, string, error) {
	bucket := "aluminum-falcon"
	gameIDBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", "", err
	}
	gameID := string(gameIDBytes)
	body := newGame()

	sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	uploader := s3manager.NewUploader(sess)
	_, uploadErr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("games/" + string(gameID) + ".json"),
		Body:   strings.NewReader(body),
	})
	if uploadErr != nil {
		return "", "", uploadErr
	}

	return gameID, body, nil
}

func newGame() string {
	return `{
	"status": "AWAITING_SHIPS",
	"player_1": {},
	"player_2": {}
}`
}

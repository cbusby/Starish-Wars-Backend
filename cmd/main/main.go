package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cbusby/Starish-Wars-Backend/internal/swb"
	"github.com/cbusby/Starish-Wars-Backend/internal/swb/persistence"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	persister := persistence.AWSS3Persister{}
	switch req.HTTPMethod {
	case "GET":
		gameID := req.PathParameters["gameID"]
		contents, err := swb.Read(persister, gameID)
		if err != nil {
			return serverError(err, "Could not get content for "+gameID)
		}
		return createGetGameResponse(gameID, contents)
	case "POST":
		gameID, body, err := swb.Create(persister)
		if err != nil {
			return serverError(err, "Could not create game")
		}
		return createNewGameResponse(gameID, body)
	default:
		return clientError(http.StatusMethodNotAllowed, "cannot respond to method "+req.HTTPMethod)
	}
}

func createNewGameResponse(gameID string, body string) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Location"] = gameID
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       body,
		Headers:    headers,
	}, nil
}

func createGetGameResponse(gameID string, body string) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Location"] = gameID
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
		Headers:    headers,
	}, nil
}

func serverError(err error, message string) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       message,
	}, nil
}

func clientError(status int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       message,
	}, nil
}

func main() {
	lambda.Start(router)
}

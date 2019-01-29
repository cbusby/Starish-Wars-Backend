package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
			if strings.HasPrefix(err.Error(), "Not found") {
				return notFoundError(err)
			}
			return serverError(err, fmt.Sprintf("Could not get content for %s", gameID))

		}
		return createGameResponse(http.StatusOK, gameID, contents)
	case "POST":
		gameID, body, err := swb.Create(persister)

		if err != nil {
			return serverError(err, "Could not create game")
		}
		return createGameResponse(http.StatusCreated, gameID, body)
	case "PUT":
		gameID := req.PathParameters["gameID"]
		game := req.Body
		newGame, err := swb.Update(persister, gameID, game)
		if err != nil {
			return serverError(err, "Could not update game")
		}
		return createGameResponse(http.StatusOK, gameID, newGame)
	default:
		return clientError(http.StatusMethodNotAllowed, "cannot respond to method "+req.HTTPMethod)
	}
}

func createGameResponse(responseCode int, gameID string, body string) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Location"] = gameID
	headers["Access-Control-Allow-Origin"] = "*"
	return events.APIGatewayProxyResponse{
		StatusCode: responseCode,
		Body:       body,
		Headers:    headers,
	}, nil
}

func notFoundError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())
	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       err.Error(),
		Headers:    headers,
	}, nil
}

func serverError(err error, message string) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())
	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       message,
		Headers:    headers,
	}, nil
}

func clientError(status int, message string) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       message,
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(router)
}

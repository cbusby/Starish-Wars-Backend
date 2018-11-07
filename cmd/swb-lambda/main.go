package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Starish-Wars-Backend/internal/swb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func router(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	switch req.HTTPMethod {
	case "POST":
		gameID, body, err := swb.Create()
		if err != nil {
			return serverError(err, "Could not create game")
		}
		return createNewGameResponse(gameID, body)
	default:
		return clientError(http.StatusMethodNotAllowed, "cannot respond to method "+req.HTTPMethod)
	}
}

func createNewGameResponse(gameID string, body string) events.APIGatewayProxyResponse {
	headers := make(map[string]string)
	headers["Location"] = gameID
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       body,
		Headers:    headers,
	}
}

func serverError(err error, message string) events.APIGatewayProxyResponse {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       message,
	}
}

func clientError(status int, message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       message,
	}
}

func main() {
	lambda.Start(router)
}

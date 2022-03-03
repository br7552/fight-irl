package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func errorResponse(status int,
	message interface{}) (events.APIGatewayProxyResponse, error) {

	return writeJSON(status, envelope{"error": message})
}

func serverErrorResponse(err error) (events.APIGatewayProxyResponse,
	error) {

	return events.APIGatewayProxyResponse{}, err
}

func badRequestResponse(err error) (events.APIGatewayProxyResponse, error) {
	return errorResponse(http.StatusBadRequest, err.Error())
}

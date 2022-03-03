package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type envelope map[string]interface{}

func writeJSON(status int,
	data envelope) (events.APIGatewayProxyResponse, error) {

	js, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	js = append(js, '\n')

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    headers,
		Body:       string(js),
	}, nil
}

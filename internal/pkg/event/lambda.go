package event

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type Request struct {
	Key       string `json:"key"`
	AccountId uint   `json:"account_id"`
}

func LambdaStart() {
	lambda.Start(HandleRequest)
}

func HandleRequest(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if err := handler(event.Body); err != nil {
		return response(http.StatusInternalServerError, err)
	}
	return response(http.StatusOK, nil)
}

func response(statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	var body string

	if err != nil {
		body = err.Error()
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "Content-Type",
			"Access-Control-Allow-Methods":     "POST",
			"Access-Control-Allow-Credentials": "true",
		},
	}, err
}

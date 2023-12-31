package event

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"report-transaction/internal/app/env"
	"report-transaction/internal/app/tools/decode"
)

type Request struct {
	Filename  string `json:"filename"`
	AccountId uint   `json:"account_id"`
}

func Lambda() {
	lambda.Start(HandleRequest)
}

func HandleRequest(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	request, err := decode.DeserializeJsonString[Request](event.Body)
	if err != nil {
		return response(http.StatusBadRequest, err)
	}

	err = handler(request.Filename, request.AccountId)
	if err != nil {
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
			"Content-Type":                     "text/plain",
			"Access-Control-Allow-Origin":      env.CorsOrigin,
			"Access-Control-Allow-Headers":     "Content-Type",
			"Access-Control-Allow-Methods":     env.AllowedMethods,
			"Access-Control-Allow-Credentials": "true",
		},
	}, err
}

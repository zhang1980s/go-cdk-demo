package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

type MyResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (MyResponse, error) {
	logrus.Infof("request: %v", request)

	return MyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: fmt.Sprintf("Hello, CDK, you've hit %s\n", request.Path),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

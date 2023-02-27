package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

func handler(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("request: %+v\n", event)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	apigwClient := apigatewaymanagementapi.NewFromConfig(cfg)

	err = sendMessage(apigwClient, event.RequestContext.ConnectionID, "Hello, CDK!!!! -1")

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: fmt.Sprintf("You've hit %s\n", event.RequestContext.RouteKey),
	}

	return response, nil
}

func sendMessage(client *apigatewaymanagementapi.Client, connectionID string, message string) error {
	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: &connectionID,
		Data:         []byte(message),
	}

	_, err := client.PostToConnection(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}

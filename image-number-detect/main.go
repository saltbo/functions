package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// APIGatewayEventHandler
func APIGatewayEventHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	image, ok := request.QueryStringParameters["image"]
	if !ok {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("param miss")
	}

	result, err := detect(image)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       result,
	}, nil
}

func main() {
	// 默认使用Lambda驱动，也可以换成其他的
	lambda.Start(APIGatewayEventHandler)
}

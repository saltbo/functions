package main

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	TGToken = os.Getenv("TG_TOKEN")
)

func AskQuestion(chatID int64, question string) error {
	bapi, err := tgbotapi.NewBotAPI(TGToken)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, question)
	if _, err := bapi.Send(msg); err != nil {
		return err
	}

	return nil
}

// HandleRequest 一个提醒事项接口，由外部定时器触发，触发后去发送一个提示/询问
func APIGatewayEventHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	chatID, err := strconv.ParseInt(request.QueryStringParameters["chatID"], 10, 64)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	if err := AskQuestion(chatID, request.Body); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	// 默认使用Lambda驱动，也可以换成其他的
	lambda.Start(APIGatewayEventHandler)
}

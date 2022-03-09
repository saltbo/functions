package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

var (
	TGToken = os.Getenv("TG_TOKEN")
	GHToken = os.Getenv("GH_TOKEN")
)

func SaveToGit(text string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GHToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	shaBytes := sha1.Sum([]byte(text))
	path := time.Now().Format("2006/0102.md")
	_, _, err := client.Repositories.CreateFile(ctx, "saltbo", "whynocode", path, &github.RepositoryContentFileOptions{
		Message: github.String("feat: add new file from tg"),
		Content: []byte(text),
		SHA:     github.String(hex.EncodeToString(shaBytes[:])),
		Branch:  github.String("master"),
		Author:  &github.CommitAuthor{Name: github.String("Ambor"), Email: github.String("saltbo@foxmail.com")},
	})
	return err
}

func ReplyAnswer(chatID int64, answer string) error {
	bapi, err := tgbotapi.NewBotAPI(TGToken)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, answer)
	if _, err := bapi.Send(msg); err != nil {
		return err
	}

	return nil
}

// HandleRequest 一个提醒事项接口，由外部定时器触发，触发后去发送一个提示/询问
func APIGatewayEventHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	update := new(tgbotapi.Update)
	if err := json.Unmarshal([]byte(request.Body), update); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	// todo 是否机器人主动发起的对话，如果是则根据机器人的问题处理人类的回复；如果不是则进入其他命令处理逻辑

	answer := "记录成功"
	err := SaveToGit(update.Message.Text)
	if err != nil {
		answer = fmt.Sprintf("记录失败：%s", err)
	}

	if err := ReplyAnswer(update.Message.Chat.ID, answer); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	fmt.Println(time.Now())
	// 默认使用Lambda驱动，也可以换成其他的
	lambda.Start(APIGatewayEventHandler)
}

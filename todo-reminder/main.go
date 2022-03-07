package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

// HandleRequest 一个提醒事项接口，由外部定时器触发，触发后去发送一个提示/询问
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	fmt.Println(ctx, name)

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	// 默认使用Lambda驱动，也可以换成其他的
	lambda.Start(HandleRequest)
}

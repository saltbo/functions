package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// imageTxt := r.PostFormValue("src")
	// imgData, err := base64.StdEncoding.DecodeString(imageTxt)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// sseract := gosseract.NewClient()
	// defer sseract.Close()
	//
	// if err := sseract.SetImageFromBytes(imgData); err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	//
	// text, err := sseract.Text()
	// if err != nil {
	// 	c.AbortWithError(http.StatusInternalServerError, err)
	// }
	//
	// c.JSON(http.StatusOK, gin.H{
	// 	"status": 200,
	// 	"msg":    text,
	// })

	msg := "Hello, world!\n"
	w.Write([]byte(msg))
}

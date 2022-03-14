package main

import (
	"encoding/base64"

	"github.com/otiai10/gosseract/v2"
)

func detect(base64Image string) (string, error) {
	imgData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", err
	}

	sseract := gosseract.NewClient()
	defer sseract.Close()

	if err := sseract.SetImageFromBytes(imgData); err != nil {
		return "", err
	}

	return sseract.Text()
}

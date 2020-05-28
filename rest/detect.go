package rest

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

func DetectNumber(c *gin.Context) {
	imageTxt := c.PostForm("src")
	imgData, err := base64.StdEncoding.DecodeString(imageTxt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sseract := gosseract.NewClient()
	defer sseract.Close()

	if err := sseract.SetImageFromBytes(imgData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	text, err := sseract.Text()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    text,
	})
}

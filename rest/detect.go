package rest

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

var sseract *gosseract.Client

func init() {
	sseract = gosseract.NewClient()
}

func DetectNumber(c *gin.Context) {
	imageTxt := c.PostForm("image")
	imgData, err := base64.StdEncoding.DecodeString(imageTxt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := sseract.SetImageFromBytes(imgData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	text, err := sseract.Text()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"text": text,
	})
}

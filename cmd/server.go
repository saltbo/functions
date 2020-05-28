package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"gofc/rest"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/detect/number", rest.DetectNumber)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

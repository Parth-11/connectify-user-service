package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	if err := router.Run("8080"); err != nil {
		fmt.Println("Failed to start server on the port 8080", err)
	}
}

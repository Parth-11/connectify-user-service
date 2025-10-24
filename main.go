package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	router := gin.Default()

	con, err := pgx.Connect(context.Background(), "")

	if err != nil {
		fmt.Println("Unable to connect to database", err)
		os.Exit(1)
	}
	defer con.Close(context.Background())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	if err := router.Run("8080"); err != nil {
		fmt.Println("Failed to start server on the port 8080", err)
	}
}

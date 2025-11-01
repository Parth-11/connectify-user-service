package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Parth-11/connectify-user-service/models"
	"github.com/Parth-11/connectify-user-service/repository"
	"github.com/Parth-11/connectify-user-service/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}
	user.Password = hashedPwd

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repository.CreateUser(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":    user.UserID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login LoginInput

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email := login.Email
		password := login.Password

		c.JSON(200, gin.H{"EMAIL": email,
			"PASSWORD": password})
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Inside returned handler!"})
	}
}

func AdminDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Inside returned handler!"})
	}
}

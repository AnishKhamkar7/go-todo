package controllers

import "github.com/gin-gonic/gin"

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

}

func LoginUser() gin.HandlerFunc {

}

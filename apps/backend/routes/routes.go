package routes

import (
	"github.com/AnishKhamkar7/todo-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	v1 := r.Group("/v1/users")

	v1.POST("/login", controllers.RegisterUser())
	v1.POST("/register", controllers.LoginUser())

}

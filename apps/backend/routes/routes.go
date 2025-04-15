package routes

import (
	"github.com/AnishKhamkar7/todo-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Inside returned handler!"})
	})

	v1 := r.Group("/v1/users")

	v1.POST("/register", controllers.RegisterUser())

	v1.POST("/login", controllers.LoginUser())
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
	}))

	authorized.GET("/dashboard", controllers.AdminDashboard())
}

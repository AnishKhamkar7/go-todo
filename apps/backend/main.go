package main

import (
	"fmt"
	"net/http"

	"github.com/AnishKhamkar7/todo-api/db"
	"github.com/AnishKhamkar7/todo-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	db.DbConfig()
}

func main() {

	router := gin.Default()

	routes.SetupRoutes(router)

	server := &http.Server{
		Addr:    ":8070",
		Handler: router,
	}

	fmt.Println("Server Running on port 8070")
	server.ListenAndServe()
}

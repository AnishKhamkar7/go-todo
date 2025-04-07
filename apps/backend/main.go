package main

import (
	"fmt"
	"net/http"

	"github.com/AnishKhamkar7/todo-api/db"
	"github.com/gin-gonic/gin"
)

func main() {

	db.DbConfig()

	router := gin.Default()

	server := &http.Server{
		Addr:    ":8070",
		Handler: router,
	}

	fmt.Println("Server Running on port 8070")
	server.ListenAndServe()
}

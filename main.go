package main

import (
	"github.com/Today-budget-app/backend/domain/controllers"
	"github.com/Today-budget-app/backend/infra/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.UserCreate)
	r.POST("/signin", controllers.UserFetch)
	r.POST("/transactions", controllers.TransactionCreate)
	r.GET("/transactions", controllers.TransactionFetch)
	r.GET("/transactions/PLACEHOLDER")
	r.DELETE("/transactions/PLACEHOLDER")
	r.PUT("/transactions/PLACEHOLDER")

	r.Run()
}

//25min

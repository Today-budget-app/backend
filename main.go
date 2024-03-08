package main

import (
	"github.com/Today-budget-app/backend/controllers"
	"github.com/Today-budget-app/backend/initializers"
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

	r.Run()
}

//25min

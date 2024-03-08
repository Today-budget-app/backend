package main

import (
	"github.com/Today-budget-app/backend/initializers"
	"github.com/Today-budget-app/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}) //go run migrate/migrate.go
}

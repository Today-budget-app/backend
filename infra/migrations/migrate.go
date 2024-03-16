package main

import (
	"github.com/Today-budget-app/backend/domain/models"
	"github.com/Today-budget-app/backend/infra/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})        //go run migrate/migrate.go
	initializers.DB.AutoMigrate(&models.Transaction{}) //go run migrate/migrate.go
}

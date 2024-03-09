package controllers

import (
	"time"

	"github.com/Today-budget-app/backend/initializers"
	"github.com/Today-budget-app/backend/models"
	"github.com/gin-gonic/gin"
)

type transCtrl struct {
	Amount      float32
	TranType    string
	TranDate    time.Time
	Description string
}

func TransactionCreate(c *gin.Context) {
	//get date from request body
	//create user
	//return it
	//GORM create guide

	var transac transCtrl

	c.Bind(&transac)

	transaction := models.Transaction{Amount: transac.Amount, TranType: transac.TranType, TransactionDate: transac.TranDate, Description: transac.Description}

	result := initializers.DB.Create(&transaction)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"transaction": transaction,
	})
}

func TransactionFetch(c *gin.Context) {
	//get user
	var transactions []models.Transaction

	initializers.DB.Find(&transactions)

	//respond
	c.JSON(200, gin.H{
		"transaction": transactions,
	})
}

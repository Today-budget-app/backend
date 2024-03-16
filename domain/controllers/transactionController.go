package controllers

import (
	"time"

	"github.com/Today-budget-app/backend/domain/models"
	"github.com/Today-budget-app/backend/infra/initializers"
	"github.com/gin-gonic/gin"
)

type transCtrl struct {
	Amount      float32
	TranType    string
	TranDate    time.Time
	Description string
}

func TransactionCreate(c *gin.Context) {

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

	var transactions []models.Transaction

	initializers.DB.Find(&transactions)

	c.JSON(200, gin.H{
		"transaction": transactions,
	})
}

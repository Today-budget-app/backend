package controllers

import (
	"log"
	"time"

	"github.com/Today-budget-app/backend/domain/managers"
	"github.com/Today-budget-app/backend/domain/models"
	"github.com/gin-gonic/gin"
)

type transCtrl struct {
	Amount      float32
	TranType    string
	TranDate    time.Time
	Description string
}

type TransactionsController struct {
	TransactionsManager managers.TransactionsManager
}

func NewTransactionsController(transactionsManager *managers.TransactionsManager) *TransactionsController {
	return &TransactionsController{TransactionsManager: *transactionsManager}
}

func (controller *TransactionsController) Create(c *gin.Context) {
	var transactionData models.Transaction
	err := c.BindJSON(&transactionData)
	if err != nil {
		log.Printf("Error binding transaction data: %v", err)
	}

	transaction, err := controller.TransactionsManager.Create(&transactionData)
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		c.JSON(400, gin.H{"error": "Error creating transaction"})
		return
	}

	c.JSON(200, gin.H{"transaction": transaction})
}

func (controller *TransactionsController) List(c *gin.Context) {
	transactions, err := controller.TransactionsManager.List()
	if err != nil {
		log.Printf("Error listing transactions: %v", err)
		c.JSON(400, gin.H{"error": "Error listing transactions"})
		return
	}
	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}

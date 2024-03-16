package managers

import (
	"github.com/Today-budget-app/backend/domain/models"
	"github.com/Today-budget-app/backend/infra/initializers"
)

type TransactionsManager struct{}

func NewTransactionsManager() *TransactionsManager {
	return &TransactionsManager{}
}

func (m *TransactionsManager) Create(transaction *models.Transaction) (models.Transaction, error) {
	result := initializers.DB.Create(&transaction)
	if result.Error != nil {
		return models.Transaction{}, result.Error
	}
	return *transaction, nil
}

func (m *TransactionsManager) List() (*[]models.Transaction, error) {
	var transactions []models.Transaction
	result := initializers.DB.Find(&transactions)
	if result.Error != nil {
		return &[]models.Transaction{}, result.Error
	}
	return &transactions, nil
}

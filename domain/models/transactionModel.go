package models

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model already includes ID, CreatedAt, UpdatedAt, DeletedAt, Name(String)

type Transaction struct {
	gorm.Model
	Amount          float32
	TransactionDate time.Time
	TranType        string
	Description     string
}

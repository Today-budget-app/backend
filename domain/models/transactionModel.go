package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model already includes ID, CreatedAt, UpdatedAt, DeletedAt, Name(String)

type Transaction struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Amount          float32
	TransactionDate time.Time
	TranType        string
	Description     string
}

package models

import (
	"time"
	//"gorm.io/gorm"
)

// gorm.Model already includes ID, CreatedAt, UpdatedAt, DeletedAt, Name(String)
type User struct {
	//gorm.Model
	ID        uint
	Username  string
	Email     string
	CreatedAt time.Time
}

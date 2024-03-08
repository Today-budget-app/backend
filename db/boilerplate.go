// https://gorm.io/docs/connecting_to_the_database.html
package main

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PLACEHOLDER struct {
	gorm.Model
	Place  string
	Holder uint
}

func main() {
	sqlDB, err := sql.Open("pgx", "mydb_dsn")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&PLACEHOLDER{})

	// Create
	db.Create(&PLACEHOLDER{Place: "D42", Holder: 100})

	// Read
	var PLACEHOLDER PLACEHOLDER
	db.First(&PLACEHOLDER, 1)                  // find PLACEHOLDER with integer primary key
	db.First(&PLACEHOLDER, "Place = ?", "D42") // find PLACEHOLDER with Place D42

	// Update - update PLACEHOLDER's Holder to 200
	db.Model(&PLACEHOLDER).Update("Holder", 200)
	// Update - update multiple fields
	db.Model(&PLACEHOLDER).Updates(PLACEHOLDER{Holder: 200, Place: "F42"}) // non-zero fields
	db.Model(&PLACEHOLDER).Updates(map[string]interface{}{"Holder": 200, "Place": "F42"})

	// Delete - delete PLACEHOLDER
	db.Delete(&PLACEHOLDER, 1)
}

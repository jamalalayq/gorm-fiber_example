package database

import (
	"fmt"
	"gorm_fiber/book"
)

// Migrate build migrations
func Migrate() {
	DB.AutoMigrate(
		&book.Book{},
	)
	fmt.Println("Database migrated..")
}

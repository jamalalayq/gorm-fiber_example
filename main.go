package main

import (
	"fmt"
	"gorm_fiber/database"
	"gorm_fiber/routes"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func initializeDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection succesfully opened.")

	database.Migrate()
}

func main() {
	app := fiber.New()

	initializeDatabase()
	defer database.DB.Close()

	routes.RegisterBooksRoutes(app, database.DB)

	app.Listen(3000)
}

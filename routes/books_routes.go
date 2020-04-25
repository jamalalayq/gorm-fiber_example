package routes

import (
	"gorm_fiber/book"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// RegisterBooksRoutes set books routes
func RegisterBooksRoutes(app *fiber.App, db *gorm.DB) {
	group := app.Group("/api/v1")
	controller := book.New(app, group, db)
	controller.GetBooks()
	controller.GetBook()
	controller.AddBook()
	controller.DeleteBook()
}
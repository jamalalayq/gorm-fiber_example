package book

import (
	"gorm_fiber/core"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// IBook : interface
type IBook interface {
	GetBooks()
	GetBook()
	AddBook()
	DeleteBook()
}

type controller struct {
	app   *fiber.App
	group *fiber.Group
	db    *gorm.DB
}

// New : create new instance from book.
func New(app *fiber.App, group *fiber.Group, db *gorm.DB) IBook {
	c := controller{app: app, group: group, db: db}
	return c
}

func (controller controller) createNewGroup() *fiber.Group {
	return controller.group.Group("/book")
}

// GetBooks : Fetch all books.
func (controller controller) GetBooks() {
	group := controller.createNewGroup()
	group.Get("/", func(c *fiber.Ctx) {
		booksRepo := BooksRepository{DB: controller.db}
		books := booksRepo.GetAll()
		response := core.Response{Status: "OK", Code: 200, Data: books}
		c.Status(200).JSON(response)
	})
}

// GetBook : Get one book.
func (controller controller) GetBook() {
	group := controller.createNewGroup()
	group.Get("/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		booksRepo := BooksRepository{DB: controller.db}
		book, err := booksRepo.GetOne(id)
		response := core.Response{Status: "OK", Code: 200}
		if err != nil {
			response.Err = err.Error()
			response.Status = "NOTOK"
			c.Status(200).JSON(response)
			return
		}
		response.Data = book
		c.Status(200).JSON(response)
	})
}

// AddBook : Save new book in database.
func (controller controller) AddBook() {
	group := controller.createNewGroup()
	group.Post("/", func(c *fiber.Ctx) {
		booksRepo := BooksRepository{DB: controller.db}
		book := new(Book)
		response := core.Response{Status: "OK", Code: 200}
		if err := c.BodyParser(book); err != nil {
			response.Err = err.Error()
			c.Status(200).JSON(response)
			return
		}
		savedBook := booksRepo.Save(*book)
		response.Data = savedBook
		c.Status(201).JSON(response)
	})
}

// DeleteBook : Remove book from database.
func (controller controller) DeleteBook() {
	group := controller.createNewGroup()
	group.Delete("/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		booksRepo := BooksRepository{DB: controller.db}
		err := booksRepo.Delete(id)
		response := core.Response{Status: "OK", Code: 200}
		if err != nil {
			response.Status = "NOTOK"
			response.Err = err.Error()
			c.Status(200).JSON(response)
			return
		}
		c.Status(200).JSON(response)
	})
}

package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)


type booksRoutes struct {
	booksController controllers.BooksController
}

func NewBooksRoutes(booksController controllers.BooksController) Routes{
	return &booksRoutes{booksController}
}


func (r *booksRoutes) Install(app *fiber.App) {
	app.Post("/books", r.booksController.InsertBook)
	app.Get("/books", r.booksController.GetBooks)
	app.Get("/books/:id", r.booksController.GetBook)
	app.Put("/books/:id", r.booksController.PutBook)
	app.Delete("/books/:id", r.booksController.DeleteBook)
}

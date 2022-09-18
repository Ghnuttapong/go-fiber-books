package controllers

import (
	"fiber-mongo-api/models"
	"fiber-mongo-api/repository"
	"fiber-mongo-api/util"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BooksController interface {
	GetBook(ctx *fiber.Ctx) error
	GetBooks(ctx *fiber.Ctx) error
	PutBook(ctx *fiber.Ctx) error
	InsertBook(ctx *fiber.Ctx) error
	DeleteBook(ctx *fiber.Ctx) error
}

type booksController struct {
	booksRepo repository.BooksRepository
}

func NewBooksController(booksRepo repository.BooksRepository) BooksController {
	return &booksController{booksRepo}
}

func (c *booksController) GetBook(ctx *fiber.Ctx) error {
	book, err := c.booksRepo.GetById(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(book)
}

func (c *booksController) GetBooks(ctx *fiber.Ctx) error {
	books, err := c.booksRepo.GetAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(books)
}

func (c *booksController) InsertBook(ctx *fiber.Ctx) error {
	var newBook models.Book
	err := ctx.BodyParser(&newBook)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}	
	exists, err := c.booksRepo.GetName(newBook.Name)
		
	if err == mgo.ErrNotFound {
		newBook.CreatedAt = time.Now()
		newBook.UpdatedAt = newBook.CreatedAt
		newBook.Id = bson.NewObjectId()
		err = c.booksRepo.Save(&newBook)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
		return ctx.Status(http.StatusCreated).JSON(newBook)
	}
	if exists != nil {
		err = util.ErrNameBookAlreadyExists
	}
	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
}

func (c *booksController) PutBook(ctx *fiber.Ctx) error {
	var update models.Book
	err := ctx.BodyParser(&update)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	// exists, err := c.booksRepo.GetName(update.Name)
	// if err == mgo.ErrNotFound {
		book, err := c.booksRepo.GetById(ctx.Params("id"))
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
		book.Name = update.Name
		book.Author = update.Author
		book.Price = update.Price
		book.UpdatedAt = time.Now()
		err = c.booksRepo.Update(book)
		if err != nil {
			return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
		}
		return ctx.Status(http.StatusOK).JSON(book)
	// }
	// if exists != nil {
	// 	err = util.ErrNameBookAlreadyExists
	// }
	// return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
}

func (c *booksController) DeleteBook(ctx *fiber.Ctx) error {
	err := c.booksRepo.Delete(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(err))
	}
	ctx.Set("Entity", ctx.Params("id"))
	return ctx.SendStatus(http.StatusNoContent)
} 
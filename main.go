package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/controllers"
	"fiber-mongo-api/repository"
	"fiber-mongo-api/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	conn := configs.NewConnection()
	defer conn.Close()

	app := fiber.New()
	app.Use(cors.New())	
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error{
		return c.Status(http.StatusOK).JSON(fiber.Map{"message" : "Hello World.."})
	})

	bookRepo := repository.NewBooksRepository(conn)
	booksController := controllers.NewBooksController(bookRepo)
	bookRoutes := routes.NewBooksRoutes(booksController)
	bookRoutes.Install(app)

	log.Fatal(app.Listen(":8080"))
}

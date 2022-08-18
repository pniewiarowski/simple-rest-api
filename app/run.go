package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/pniewiarowski/simple-rest-api/router"
)

func Run(port int) {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	auth := api.Group("/auth")
	book := api.Group("/book")
	author := api.Group("/author")

	router.SetupAuth(auth)
	router.SetupBook(book)
	router.SetupAuthor(author)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

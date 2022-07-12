package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/router"
	"log"
)

func Run(port int) {
	app := fiber.New()
	api := app.Group("/api")

	router.SetupBook(api)
	router.SetupAuthor(api)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

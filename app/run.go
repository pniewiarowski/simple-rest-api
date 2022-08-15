package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pniewiarowski/simple-rest-api/router"
)

func Run(port int) {
	app := fiber.New()
	api := app.Group("/api")

	app.Use(logger.New())
	app.Use(csrf.New())

	router.SetupBook(api)
	router.SetupAuthor(api)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

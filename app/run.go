package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/router"
	"log"
)

func Run(port int) {
	app := fiber.New()

	router.SetupBook(app)

	log.Fatal(
		app.Listen(
			fmt.Sprintf(":%d", port),
		),
	)
}

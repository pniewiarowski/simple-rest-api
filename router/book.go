package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pniewiarowski/simple-rest-api/controller"
	"github.com/pniewiarowski/simple-rest-api/middleware"
)

func SetupBook(router fiber.Router) {
	router.Get("/", middleware.Protected(), controller.GetAllBook)
	router.Get("/:id", middleware.Protected(), controller.GetBook)
	router.Post("/", middleware.Protected(), controller.CreateBook)
	router.Put("/:id", middleware.Protected(), controller.UpdateBook)
	router.Delete("/:id", middleware.Protected(), controller.DeleteBook)
}

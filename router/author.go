package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pniewiarowski/simple-rest-api/controller"
	"github.com/pniewiarowski/simple-rest-api/middleware"
)

func SetupAuthor(router fiber.Router) {
	router.Get("/", middleware.Protected(), controller.GetAllAuthor)
	router.Get("/:id", middleware.Protected(), controller.GetAuthor)
	router.Post("/", middleware.Protected(), controller.CreateAuthor)
	router.Put("/:id", middleware.Protected(), controller.UpdateAuthor)
	router.Delete("/:id", middleware.Protected(), controller.DeleteAuthor)
}

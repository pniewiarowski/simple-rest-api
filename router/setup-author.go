package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/controller"
)

func SetupAuthor(router fiber.Router) {
	router.Get("/author", controller.GetAllAuthor)
	router.Get("/author/:id", controller.GetAuthor)
	router.Post("/author", controller.CreateAuthor)
	router.Put("/author/:id", controller.UpdateAuthor)
	router.Delete("/author/:id", controller.DeleteAuthor)
}

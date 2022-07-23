package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/controller"
)

func SetupBook(router fiber.Router) {
	router.Get("/book", controller.GetAllBook)
	router.Get("/book/:id", controller.GetBook)
	router.Post("/book", controller.CreateBook)
	router.Put("/book/:id", controller.UpdateBook)
	router.Delete("/book/:id", controller.DeleteBook)
}

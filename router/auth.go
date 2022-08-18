package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pniewiarowski/simple-rest-api/controller"
)

func SetupAuth(router fiber.Router) {
	router.Post("/login", controller.Login)
	router.Post("/register", controller.Register)
}

package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/controller"
)

const base = "/api/book"

func SetupBook(app *fiber.App) {
	app.Get(base, controller.GetAllBook)
	app.Get(fmt.Sprintf("%s/:id", base), controller.GetBook)
	app.Post(base, controller.CreateBook)
	app.Put(fmt.Sprintf("%s/:id", base), controller.UpdateBook)
	app.Delete(fmt.Sprintf("%s/:id", base), controller.DeleteBook)
}

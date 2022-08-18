package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pniewiarowski/simple-rest-api/models"
)

func GetAuthor(ctx *fiber.Ctx) error {
	author, err := models.GetAuthorByID(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"author": author})
}

func GetAllAuthor(ctx *fiber.Ctx) error {
	authors, err := models.GetAllAuthors()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"authors": authors})
}

func CreateAuthor(ctx *fiber.Ctx) error {
	author := new(models.Author)
	if err := ctx.BodyParser(author); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	author, err := models.CreateAuthor(author)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"author": author})
}

func UpdateAuthor(ctx *fiber.Ctx) error {
	updatedAuthor := new(models.Author)
	if err := ctx.BodyParser(updatedAuthor); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	author, err := models.GetAuthorByID(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	updatedAuthor, err = models.UpdateAuthor(&author, updatedAuthor)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"author": updatedAuthor})
}

func DeleteAuthor(ctx *fiber.Ctx) error {
	authors, err := models.DeleteAuthor(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"success": false})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"authors": authors})
}

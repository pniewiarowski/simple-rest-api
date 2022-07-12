package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/database"
	"github.com/pniewiarowski/simple-rest-api/models"
)

func GetAuthor(ctx *fiber.Ctx) error {
	var author models.Author

	database.DataBase.First(&author, ctx.Params("id"))

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"author":  author,
	})
}

func GetAllAuthor(ctx *fiber.Ctx) error {
	var authors []models.Author

	database.DataBase.First(&authors)

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"authors": authors,
	})
}

func CreateAuthor(ctx *fiber.Ctx) error {
	author := new(models.Author)

	if err := ctx.BodyParser(&author); err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	database.DataBase.Create(author)

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"author":  author,
	})
}

func UpdateAuthor(ctx *fiber.Ctx) error {
	var author models.Author
	updatedAuthor := new(models.Author)

	if err := ctx.BodyParser(&updatedAuthor); err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	database.DataBase.Model(author).Updates(updatedAuthor)
	updatedAuthor.ID = author.ID

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"author":  updatedAuthor,
	})
}

func DeleteAuthor(ctx *fiber.Ctx) error {
	var author models.Author
	database.DataBase.Delete(&author, ctx.Params("id"))

	var authors []models.Author
	database.DataBase.Find(&authors)

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"author":  authors,
	})
}

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/models"
)

func GetBook(ctx *fiber.Ctx) error {
	book, err := models.GetBookByID(ctx.Params("id"))

	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"book":    book,
	})
}

func GetAllBook(ctx *fiber.Ctx) error {
	books, err := models.GetAllBooks()

	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"books":   books,
	})
}

func CreateBook(ctx *fiber.Ctx) error {
	book := new(models.Book)
	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	book, err := models.CreateBook(book)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"book":    book,
	})
}

func UpdateBook(ctx *fiber.Ctx) error {
	updatedBook := new(models.Book)
	if err := ctx.BodyParser(updatedBook); err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	book, err := models.GetBookByID(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	updatedBook, err = models.UpdateBook(&book, updatedBook)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"book":    updatedBook,
	})
}

func DeleteBook(ctx *fiber.Ctx) error {
	books, err := models.DeleteBook(ctx.Params("id"))

	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"books":   books,
	})
}

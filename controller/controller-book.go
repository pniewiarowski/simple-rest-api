package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pniewiarowski/simple-rest-api/database"
	"github.com/pniewiarowski/simple-rest-api/models"
	"log"
)

func GetBook(ctx *fiber.Ctx) error {
	var book models.Book

	database.DataBase.First(&book, ctx.Params("id"))

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"book":    book,
	})
}

func GetAllBook(ctx *fiber.Ctx) error {
	var books []models.Book

	database.DataBase.Find(&books)

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

	database.DataBase.Create(&book)

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"book":    book,
	})
}

func UpdateBook(ctx *fiber.Ctx) error {
	var book models.Book
	updatedBook := new(models.Book)

	database.DataBase.First(&book, ctx.Params("id"))

	if err := ctx.BodyParser(updatedBook); err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	database.DataBase.Model(book).Updates(updatedBook)
	updatedBook.ID = book.ID

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"book":    updatedBook,
	})
}

func DeleteBook(ctx *fiber.Ctx) error {
	var book models.Book
	database.DataBase.Delete(&book, ctx.Params("id"))

	var books []models.Book
	database.DataBase.Find(&books)
	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"books":   books,
	})
}

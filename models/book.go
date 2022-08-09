package models

import (
	"github.com/pniewiarowski/simple-rest-api/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	AuthorID    int     `json:"author-id"`
	Author      Author  `json:"author" gorm:"foreignKey:AuthorID"`
}

func GetBookByID(id string) (Book, error) {
	var book Book

	err := database.DataBase.Preload("Author").First(&book, id).Error

	return book, err
}

func GetAllBooks() ([]Book, error) {
	var books []Book

	err := database.DataBase.Preload("Author").Find(&books).Error

	return books, err
}

func CreateBook(book *Book) (*Book, error) {
	database.DataBase.Create(&book)
	err := database.DataBase.Preload("Author").Find(&book).Error

	return book, err
}

func UpdateBook(book *Book, updatedBook *Book) (*Book, error) {
	database.DataBase.Model(book).Updates(&updatedBook)

	updatedBook.ID = book.ID

	err := database.DataBase.Preload("Author").Find(&updatedBook).Error

	return updatedBook, err
}

func DeleteBook(id string) ([]Book, error) {
	var book Book
	var books []Book

	err := database.DataBase.Delete(&book, id).Error
	if err != nil {
		return []Book{}, err
	}

	err = database.DataBase.Preload("Author").Find(&books).Error

	return books, err
}

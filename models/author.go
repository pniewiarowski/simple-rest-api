package models

import (
	"gorm.io/gorm"

	"github.com/pniewiarowski/simple-rest-api/database"
)

type Author struct {
	gorm.Model
	FirstName  string `json:"firstname"`
	FamilyName string `json:"familyname"`
	Books      []Book `json:"books" gorm:"foreignKey:AuthorID"`
}

func GetAuthorByID(id string) (Author, error) {
	var author Author

	err := database.DataBase.Preload("Books").First(&author, id).Error

	return author, err
}

func GetAllAuthors() ([]Author, error) {
	var authors []Author

	err := database.DataBase.Preload("Books").Find(&authors).Error

	return authors, err
}

func CreateAuthor(author *Author) (*Author, error) {
	database.DataBase.Create(&author)
	err := database.DataBase.Preload("Books").Find(&author).Error

	return author, err
}

func UpdateAuthor(author *Author, updatedAuthor *Author) (*Author, error) {
	database.DataBase.Model(author).Updates(&updatedAuthor)

	updatedAuthor.ID = author.ID

	err := database.DataBase.Preload("Books").Find(&updatedAuthor).Error

	return updatedAuthor, err
}

func DeleteAuthor(id string) ([]Author, error) {
	var author Author
	var authors []Author

	err := database.DataBase.Delete(&author, id).Error
	if err != nil {
		return []Author{}, err
	}

	err = database.DataBase.Preload("Books").Find(&authors).Error

	return authors, err
}

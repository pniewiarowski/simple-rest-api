package models

import (
	"gorm.io/gorm"

	"github.com/pniewiarowski/simple-rest-api/database"
)

type Auth struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Enable   bool   `json:"enable"`
}

func CreateAuth(auth *Auth) error {
	return database.DataBase.Create(&auth).Error
}

func GetAuthByEmail(email string) *Auth {
	auth := new(Auth)
	database.DataBase.First(&auth, "email = ?", email)

	return auth
}

package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Author      string  `json:"author"`
}

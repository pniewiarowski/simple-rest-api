package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	AuthorID    int     `json:"author-id"`
	Author      Author  `json:"author"`
}

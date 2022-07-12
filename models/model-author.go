package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	FirstName  string `json:"first-name"`
	FamilyName string `json:"family-name"`
}

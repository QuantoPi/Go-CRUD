package models

import "gorm.io/gorm"

type Costumer struct {
	gorm.Model

	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string
	PhoneNumber string
}

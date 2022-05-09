package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Mail     string `gorm:"unique"`
	Password []byte
	Age      uint8
}

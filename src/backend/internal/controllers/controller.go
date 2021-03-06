package controllers

import "gorm.io/gorm"

type Base struct {
	DB *gorm.DB
}

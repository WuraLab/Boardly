package models

import (
	"errors"

	"gorm.io/gorm"
)

//User Struct to hold User entity. Update struct tag for binding
type User struct {
	Base
	FirstName string `gorm:"column:firstname;type:varchar(255)"  json:"firstname" binding:""`
	LastName  string `gorm:"column:lastname;type:varchar(255);not null" json:"lastname" binding:""`
	Email     string `gorm:"column:email;unique;type:varchar(255);not null" json:"email" binding:"required,email"`
	Password  string `gorm:"column:password;type:varchar(255);not null" json:"password" binding:"required"`
	IsAdmin   bool `gorm:"column:isAdmin;type:BOOLEAN" json:"isAdmin"`
}

//Create This create the user struct
func (u *User) Create(db *gorm.DB) (int, error) {
	//check user already exist
	if u.Exists(db) {
		return 1, errors.New("user exists")
	}
	result := db.Create(&u)
	return int(result.RowsAffected), result.Error
}

//GetUser
func (u *User) Exists(db *gorm.DB) bool {
	result := db.Where("email = ?", u.Email).First(u)
	return result.Error == nil
}

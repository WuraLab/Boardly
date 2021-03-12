package models

import (
	"errors"

	"gorm.io/gorm"
)

//User Struct to hold User entity. Update struct tag for binding
type User struct {
	Base
	FirstName string `gorm:"column:firstname;type:varchar(255)"  json:"firstname"`
	LastName  string `gorm:"column:lastname;type:varchar(255);not null" json:"lastname"`
	Email     string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Password  string `gorm:"column:password;not null" json:"password"`
}

//Create This create the user struct
func (u *User) Create(db *gorm.DB) (int, error) {
	//check user already exist
	if u.Exists(db) {
		return 1, errors.New("user exists")
	}
	result := db.Debug().Create(&u)
	return int(result.RowsAffected), result.Error
}

//GetUser
func (u *User) Exists(db *gorm.DB) bool {
	result := db.Where("email = ?", u.Email).First(u)
	return result.Error == nil
}

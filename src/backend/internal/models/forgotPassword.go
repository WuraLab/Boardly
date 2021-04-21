package models

type ForgotPassword struct {
	Base
	FirstName string `gorm:"column:firstname;type:varchar(255)"  json:"firstname" binding:""`
	LastName  string `gorm:"column:lastname;type:varchar(255);not null" json:"lastname" binding:""`
	Email     string `gorm:"column:email;unique;type:varchar(255);not null" json:"email" binding:"required,email"`
	Password  string `gorm:"column:password;type:varchar(255);not null" json:"password" binding:"required"`
}
package auth

import (
	"fmt"
	"regexp"
)

// type Email struct {
// 	Email     string `gorm:"column:email;unique;type:varchar(255);not null" json:"email" binding:"required,email"`
// }

func IsEmailValid(email string) bool{
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	matches := pattern.MatchString(email)

	return matches
}
package models

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
	if rows, _ := u.userExists(db); rows > 0 {
		return rows, errors.New("user exists")
	}
	if err := u.hashPassword(); err != nil {
		log.Error(err)
		return 0, err
	}
	result := db.Create(&u)
	return int(result.RowsAffected), result.Error
}

//Password hash
func (u *User) hashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return err
	}
	u.Password = string(hash)
	return nil
}

//UpdateName  updates user's name
func (u *User) UpdateName(db *gorm.DB) error {
	return db.Model(&u).Updates(User{FirstName: u.FirstName, LastName: u.LastName}).Error
}

//UpdateEmail  updates user's email
func (u *User) UpdateEmail(db *gorm.DB) error {
	return db.Model(&u).Update("email", u.Email).Error
}

//UpdatePassword  updates user's name
func (u *User) UpdatePassword(db *gorm.DB) error {
	if err := u.hashPassword(); err != nil {
		log.Error(err)
		return err
	}
	return db.Model(&u).Update("password", u.Password).Error
}

//BeforeSave  gorm hook
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if err := u.hashPassword(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

//Save save entire model
func (u *User) Save(db *gorm.DB) error {
	return db.Save(&u).Error
}

//Delete  delete user operation
func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(&u).Error
}

//GetUser
func (u *User) userExists(db *gorm.DB) (int, error) {
	result := db.Where("email = ?", u.Email).First(u)
	return int(result.RowsAffected), result.Error
}

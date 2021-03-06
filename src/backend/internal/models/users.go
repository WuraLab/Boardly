package models

import (
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User Struct to hold User entity. Update struct tag for binding
type User struct {
	Base
	FirstName string `gorm:"column:firstname;type:varchar(255)"`
	LastName  string `gorm:"column:lastname;type:varchar(255);not null"`
	Email     string `gorm:"column:email;type:varchar(255);not null" binding:"email"`
	Password  string `gorm:"column:password;not null" json:"-"`
}

//UserModel ...
type UserModel struct{}

//Create This create the user struct
func (u *User) Create(db *gorm.DB) (int, error) {
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Error(err)
		return 0, err
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

//RegisterAdmin ... Create an admin user
func (m UserModel) RegisterAdmin(user *User) (err error) {
	db := db.GetDB()
	// hash password using bcrypt
	bytePassword := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err) //Something really went wrong here...
	}

	// set password to the hashed password
	user.Password = string(hashedPassword)

	if err = db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

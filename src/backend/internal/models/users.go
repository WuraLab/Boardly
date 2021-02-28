package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//User Struct to hold User entity
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt int64 `gorm:"autoUpdateTime:nano"`
	UpdatedAt int64 `gorm:"autoUpdateTime:nano"`
}

//Create This create the user struct
func (u *User) Create(db *gorm.DB) (int, error) {
	if err := db.AutoMigrate(&User{}); err != nil {
        log.Fatal("I could not create User schema")
	}

	result := db.Create(&u)
	return int(result.RowsAffected), result.Error
}

//UpdateName  updates user's name
func (u *User) UpdateName(db *gorm.DB, newName string) error {
	result := db.Model(&u).Update("name", newName)
	return result.Error
}

//Delete  delete user operation
func (u *User) Delete(db *gorm.DB) error {
	result := db.Delete(&u)
	return result.Error
}

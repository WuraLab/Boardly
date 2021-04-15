package models

import (
	"gorm.io/gorm"
)

type Role struct {
	Base
	Name  string `gorm:"column:role;type:varchar(255)"  json:"role" binding:""`
}

//Create this create the Role struct
func (r *Role) Create(db *gorm.DB) (int, error) {
	result := db.Create(&r)
	return int(result.RowsAffected), result.Error
}

//Save this saves entire Role
func (r *Role) Save(db *gorm.DB) (int, error) {
	result := db.Save(&r)
	return int(result.RowsAffected), result.Error
}

//Delete  delete role
func (r *Role) Delete(db *gorm.DB) error {
	return db.Delete(&r).Error
}

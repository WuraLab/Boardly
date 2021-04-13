package models

import (
	"errors"
	"gorm.io/gorm"
)

// only a super admin can create other admin users with roles, 
// the created allows to recall, nil for the super-admin
type Admin struct {
	Base
	User
	Role
}

//Create this create the Role struct
func (admin *Admin) Create(db *gorm.DB) (int, error) {
	result := db.Create(&admin)
	return int(result.RowsAffected), result.Error
}

//Save this saves entire Role
func (admin *Admin) Save(db *gorm.DB) (int, error) {
	result := db.Save(&admin)
	return int(result.RowsAffected), result.Error
}

//Delete  delete role
func (admin *Admin) Delete(db *gorm.DB) error {
	return db.Delete(&admin).Error
}

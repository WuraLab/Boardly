package models

import (
	"time"
)

// User type struct/class
type User struct {
	ID           uint      `json:"id"`
	CompName     string    `json:"compName"`
	CompEmail    string    `json:"compEmail"`
	CompPassword string    `json:"compPassword"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
}

// TableName ... material.tableName
func (b *User) TableName() string {
	return "materials"
}

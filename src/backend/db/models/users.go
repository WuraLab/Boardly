package models

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     *string
	CreatedAt int64 `gorm:"autoUpdateTime:nano"`
	UpdatedAt int64 `gorm:"autoUpdateTime:nano"`
}


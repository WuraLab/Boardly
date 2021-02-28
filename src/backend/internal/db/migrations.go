package db

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	"gorm.io/gorm"
)

//Migrate To create models schema in DB
func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// User table
		{
			ID: time.Now().UTC().String(),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user")
			},
		},
		// Role table
		{
			ID: time.Now().UTC().String(),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Role{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("role")
			},
		},
	})
	return m.Migrate()
}

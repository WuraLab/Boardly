package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/wuraLab/boardly/src/backend/db/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: time.Now().UTC().String(),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user")
			},
		},
	})
	return m.Migrate()
}

package models

import (
	"fmt"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/robogeek95/ginCrud/config"
)

// RegisterAdmin ... Register admin
func RegisterAdmin(user *User) (err error) {
	user.Updated = time.Now()
	user.Created = time.Now()

	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return
}

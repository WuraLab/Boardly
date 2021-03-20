package controllers_test

import (
	"github.com/wuraLab/boardly/src/backend/internal/db"
	"fmt"
	"log"

	"github.com/wuraLab/boardly/src/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
)

func init() {
	var err error
	if Config, err = config.LoadTestConfig(); err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable", Config.Database.Host, Config.Database.Port, Config.Database.User, Config.Database.Password, Config.Database.DBName)
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	}
	if err := db.Migrate(DB); err != nil {
		log.Fatalln(err)
	}
}

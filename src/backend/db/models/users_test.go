package models_test

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wuraLab/boardly/src/backend/db/models"
	"github.com/wuraLab/boardly/src/backend/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
	helper.LoadENV()
	pool, resource, _ := helper.StartDBContainer("10")
	dbHost, dbPort := os.Getenv("DB_HOST"), os.Getenv("DB_PORT")
	dbUser, dbPassword, dbName := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err, d)
	}
	DB = d
	exitVal := m.Run()
	defer os.Exit(exitVal)
	defer helper.PurgeDBContainer(pool, resource)
	runtime.Goexit() //using this to ensure container is purged
}

func TestUser(t *testing.T) {
	u := &models.User{Name: "John", Email: "johndoe@gmail.com"}
	t.Run("Create", func(t *testing.T) {
		rows, err := u.Create(DB)
		assert.NoError(t, err, "Error create user in DB")
		assert.Equal(t, rows, 1, "Rows return is not 1")
	})
	t.Run("Update", func(t *testing.T) {
		newName := "Peter"
		assert.NoError(t, u.Updatename(DB, newName), "Error create user in DB")
	})
	t.Run("Delete", func(t *testing.T) {
		assert.NoError(t, u.Delete(DB), "Error create user in DB")
	})
}

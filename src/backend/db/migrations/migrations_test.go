package migrations_test

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/wuraLab/boardly/src/backend/db/migrations"

	"github.com/stretchr/testify/assert"

	"github.com/wuraLab/boardly/src/backend/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	helper.LoadENV()
	pool, resource, _ := helper.StartDBContainer("10")
	exitVal := m.Run()
	defer os.Exit(exitVal)
	defer helper.PurgeDBContainer(pool, resource)
	runtime.Goexit() //using this to ensure container is purged
}

func TestMigrate(t *testing.T) {
	dbHost, dbPort := os.Getenv("DB_HOST"), os.Getenv("DB_PORT")
	dbUser, dbPassword, dbName := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err, d)
	}
	assert.NoError(t, migrations.Migrate(d), "Migration was not successful")
}

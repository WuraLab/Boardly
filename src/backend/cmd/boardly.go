package main

import (
	"flag"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/config"
	"github.com/wuraLab/boardly/src/backend/internal/db"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	"github.com/wuraLab/boardly/src/backend/internal/routes"
	_ "gorm.io/driver/postgres"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	configFile := flag.String("config", ".env", "App YAML Configuration Path")
	cfg, err := config.LoadConfig(*configFile, "./src/backend/internal/config")
	if err != nil {
		log.Fatal(err)
	}

	//Start PostgreSQL database
	//Example: db.GetDB()
	db.Init()

	// Initialize the routes
	r := routes.SetupRouter()

	DB := db.GetDB()
	DB.AutoMigrate(&models.User{})
	// Start serving the application
	log.Println("serving application")
	r.Run()
	log.Print(cfg)
}

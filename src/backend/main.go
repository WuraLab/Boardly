package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/db/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	migrate := flag.Bool("migrate", false, "Migrate the database schema")
	flag.Parse()

	//check the environment else set default
	env := "dev"
	if os.Getenv("ENV") != "" {
		env = os.Getenv("ENV")
	}
	envFile := ".env." + env + ".local"
	if godotenv.Load(envFile) != nil {
		log.Fatalf("Error Loading env file %s", envFile)
	}

	//create db connection
	timezone, _ := time.Now().Zone()
	log.Println(timezone)
	dbHost, dbPort, timezone := os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), timezone
	dbUser, dbPassword, dbName := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	//Do migration schema
	if *migrate {
		if migrations.Migrate(db) != nil {
			log.Fatalln("Migration was not successful")
		} else {
			log.Println("Migration was successful")
		}
		return
	}

	//star
}

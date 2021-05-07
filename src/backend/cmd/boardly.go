package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/config"
	"github.com/wuraLab/boardly/src/backend/internal/db"
	"github.com/wuraLab/boardly/src/backend/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//  fo run ....... --migrate
func init() {
	var err error
	var Config config.Config

	if Config, err = config.LoadConfig(".env"); err != nil {
		log.Fatal(err)
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	if Config.Server.LOG_LEVEL == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	}
	if Config.Server.LOG_LEVEL == "PRODUCTION" {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	var DB *gorm.DB
	var err error
	var Config config.Config

	migrate := flag.Bool("migrate", true, "display colorized output")

	//Load configuration
	if Config, err = config.LoadConfig(".env"); err != nil {
		log.Fatal(err)
	}
	//connect the DB
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable", Config.Database.Host, Config.Database.Port, Config.Database.User, Config.Database.Password, Config.Database.DBName)
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{net
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		log.Fatalln(err)
	}

	//do migration
	if *migrate {
		if err = db.Migrate(DB); err != nil {
			log.Fatal(err)
		}
	}

	if Config.Server.ENV == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	{

		// route intialization
		r := routes.SetupRouter(DB, &Config)

		log.Infoln(Config.Server.SSL)

		if Config.Server.SSL == "TRUE" {

			SSLKeys := &struct {
				CERT string
				KEY  string
			}{}

			//Generated using sh generate-certificate.sh
			SSLKeys.CERT = "./cert/myCA.cer"
			SSLKeys.KEY = "./cert/myCA.key"

			log.Fatal(r.RunTLS(":"+Config.Server.Port, SSLKeys.CERT, SSLKeys.KEY))
		} else {
			log.Fatal(r.Run(":" + Config.Server.Port))
		}
	}
}

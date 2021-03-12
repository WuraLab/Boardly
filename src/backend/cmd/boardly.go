package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/config"
	"github.com/wuraLab/boardly/src/backend/internal/db"
	"github.com/wuraLab/boardly/src/backend/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//  fo run ....... --migrate
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
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	}

	//do migration
	if *migrate {
		db.Migrate(DB)
	}

	{

		// route intialization
		r := routes.SetupRouter(DB)

		log.Infoln(Config.Server.SSL)

		if Config.Server.ENV == "PRODUCTION" {
			gin.SetMode(gin.ReleaseMode)
		}

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

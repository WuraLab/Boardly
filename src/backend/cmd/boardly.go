package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/config"
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
	log.Print(cfg)
}

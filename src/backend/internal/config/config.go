package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:",squash"`
	Database Database `mapstructure:",squash"`
}

type Database struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

type Server struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(configFile string, paths ...string) (Config, error) {
	config := Config{}
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")

	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

func LoadTestConfig() (Config, error) {
	return LoadConfig(".test.env", os.Getenv("TEST_DIR"))
}

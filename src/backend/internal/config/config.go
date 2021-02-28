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

// type Config struct {
// 	Database struct {
// 		Host        string `yaml:"host" env:"DB_HOST" env-description:"Database host" env-upd`
// 		Port        string `yaml:"port" env:"DB_PORT" env-description:"Database port" env-upd`
// 		Username    string `yaml:"username" env:"DB_USER" env-description:"Database user name" env-upd`
// 		Password    string `env:"DB_PASSWORD" env-description:"Database user password" env-upd`
// 		Name        string `yaml:"db-name" env:"DB_NAME" env-description:"Database name" env-upd`
// 		Connections int    `yaml:"connections" env:"DB_CONNECTIONS" env-description:"Total number of database connections"  env-upd`
// 	} `yaml:"database"`
// 	Server struct {
// 		Host string `yaml:"host" env:"SRV_HOST,HOST" env-description:"Server host" env-default:"localhost"  env-upd`
// 		Port string `yaml:"port" env:"SRV_PORT,PORT" env-description:"Server port" env-default:"8080"`
// 	} `yaml:"server"`
// 	Greeting string `env:"GREETING" env-description:"Greeting phrase" env-default:"Hello!"  env-upd`
// }

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

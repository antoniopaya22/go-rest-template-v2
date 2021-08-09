package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

// GetConfiguration returns the configuration
//
// parameters:
// 		configFile - the configuration file
//
// returns:
// 		the configuration
func GetConfiguration(configFile string) Configuration {
	var configuration *Configuration

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return *configuration
}

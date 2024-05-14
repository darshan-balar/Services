package inits

import (
	"log"

	"github.com/spf13/viper"
)

var ENV string

func LoadConfig(env string) {

	ENV = env
	// Set the path to your configuration file
	viper.SetConfigFile("./config/config.yaml")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read config file: ", err)
	}

	viper.SetConfigType("yaml")
}

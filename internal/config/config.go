package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/ternakkode/packform-backend/pkg/bunclient"
)

type Configuration struct {
	APPName string                   `mapstructure:"APP_NAME"`
	APPPort string                   `mapstructure:"APP_PORT"`
	DB      bunclient.DatabaseConfig `mapstructure:",squash"`
}

var environmentVariableKeys = []string{
	"APP_NAME",
	"APP_PORT",
	"DB_HOST",
	"DB_PORT",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",
	"DB_SSL_MODE",
	"DB_URL",
	"DB_DEBUG",
	"DB_DEBUG_LEVEL",
}

var config *Configuration

func Init() *Configuration {
	viper.AutomaticEnv()

	for _, key := range environmentVariableKeys {
		viper.BindEnv(key, key)
	}

	config = new(Configuration)
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal("error while unmarshalling config: ", err)
	}

	return config
}

func GetConfig() *Configuration {
	return config
}

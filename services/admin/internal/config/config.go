package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./internal/config")

	if err := Config.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file: ", err)
	}
}

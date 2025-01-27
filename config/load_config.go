package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".") // working directory
	viper.SetConfigType("yml")
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

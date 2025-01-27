package config

import (
	"os"

	"github.com/spf13/viper"
)

func ExpandEnv() {
	mapping := func(s string) string {
		return viper.GetString(s)
	}

	for _, key := range viper.AllKeys() {
		v := viper.GetString(key)
		viper.Set(key, os.Expand(v, mapping))
	}
}

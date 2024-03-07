package config

import "github.com/spf13/viper"

func InitViper(path string) error {
	viper.AddConfigPath(path)
	return viper.ReadInConfig()
}

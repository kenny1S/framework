package config

import "github.com/spf13/viper"

func InitViper() error {
	viper.AddConfigPath("/Users/maoyuting/go/src/2108a/high-five/home/day7/user_rpc/config")
	return viper.ReadInConfig()
}

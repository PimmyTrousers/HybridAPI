package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	apikey string
	// only to be used for deprecated v1 api endpoints
	apisecret string
	server    string
}

func GetConfig() (*Config, error) {
	conf := &Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}
	conf.apikey = viper.Get("apikey").(string)
	conf.server = viper.Get("server").(string)
	conf.apisecret = viper.Get("apisecret").(string)
	return conf, nil
}

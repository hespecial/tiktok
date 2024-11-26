package config

import (
	"github.com/spf13/viper"
	"log"
)

var Conf *Config

type Config struct {
	App   App
	Mysql Mysql
}

func InitConfig() {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read in config error: ", err.Error())
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatal("Init config error: ", err.Error())
	}
}

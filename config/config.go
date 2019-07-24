package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	AppName string
	Port int
	DataBaseHost string
	DataBaseName string
}

var appConfig *AppConfig

func init() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.Parse()
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error reading config file")
	}
	appConfig = CreateConfig();
}

func GetConfig() *AppConfig {
	return appConfig
}

func CreateConfig() *AppConfig{
	config := new(AppConfig)
	config.AppName = viper.GetString("appName")
	config.Port = viper.GetInt("port")
	config.DataBaseHost = viper.GetString("mongo.host")
	config.DataBaseName = viper.GetString("mongo.database")
	return config
}
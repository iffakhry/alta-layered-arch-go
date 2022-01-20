package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//AppConfig Application configuration
type AppConfig struct {
	Port     int `json:"port" yaml:"port"`
	Database struct {
		Driver   string `json:"driver" yaml:"driver"`
		Name     string `json:"name" yaml:"name"`
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8080
	defaultConfig.Database.Driver = viper.GetString(`database.driver`)
	defaultConfig.Database.Name = viper.GetString(`database.name`)
	defaultConfig.Database.Host = viper.GetString(`database.host`)
	defaultConfig.Database.Port = viper.GetInt(`database.port`)
	defaultConfig.Database.Username = viper.GetString(`database.username`)
	defaultConfig.Database.Password = viper.GetString(`database.password`)

	// viper.SetConfigType("yaml")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		// log.Info("error to load config file, will use default value ", err)
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig
}

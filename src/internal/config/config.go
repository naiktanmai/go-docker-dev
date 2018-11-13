package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

type Constants struct {
	PORT  string
	Mongo struct {
		URL    string
		DBName string
	}
}

type Config struct {
	Constants
	Database *mgo.Database
}

func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants

	if err != nil {
		return &config, err
	}
	dbSession, err := mgo.Dial(config.Constants.Mongo.URL)
	if err != nil {
		return &config, err
	}

	config.Database = dbSession.DB(config.Constants.Mongo.DBName)
	return &config, err
}

func initViper() (Constants, error) {
	viper.SetConfigName("todo.config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Error reading config file, %s", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.SetDefault("PORT", "3001")

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}

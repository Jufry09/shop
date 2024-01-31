package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port int
	DB   DB
}

type DB struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

func GetConfig() Config {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return Config{
		Port: viper.GetInt("port"),
		DB: DB{
			Host:         viper.GetString("db.host"),
			Port:         viper.GetInt("db.port"),
			Username:     viper.GetString("db.username"),
			Password:     viper.GetString("db.password"),
			DatabaseName: viper.GetString("db.database_name"),
		},
	}

}

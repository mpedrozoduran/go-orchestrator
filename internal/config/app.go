package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Bank     BankConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Url string
}

type BankConfig struct {
	Url string
}

type AuthConfig struct {
	Username  string
	Password  string
	SecretKey string
}

func LoadConfig() Config {
	viper.SetConfigFile("./config/config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %s\n", err)
		return Config{}
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Error unmarshaling config: %s\n", err)
		return Config{}
	}
	return config
}

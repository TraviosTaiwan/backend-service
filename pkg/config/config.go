package config

import (
	"log"

	"github.com/spf13/viper"
)

var LocalConfig *Config

type Config struct {
	DBUser        string `mapstructure:"DB_USERNAME"`
	DBPass        string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBName        string `mapstructure:"DB_DATABASE"`
	DBPort        string `mapstructure:"DB_PORT"`
	Port          string `mapstructure:"PORT"`
	ApiHost       string `mapstructure:"API_HOST"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiration int    `mapstructure:"JWT_EXPIRATION"`
}

func initConfig() *Config {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}
	return config
}

func SetConfig() {
	LocalConfig = initConfig()
}

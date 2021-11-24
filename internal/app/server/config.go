package server

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	LogLevel string
}

func NewConfig(configPath string) *Config {
	pathParts := strings.Split(configPath, "/")
	filepath := strings.Join(pathParts[:len(pathParts)-1], "/")
	filename := pathParts[len(pathParts)-1]

	viper.AddConfigPath(filepath)
	viper.SetConfigName(filename)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config: %s", err.Error())
	}

	return &Config{
		Port:     viper.GetString("server.port"),
		LogLevel: viper.GetString("server.log_level"),
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		Port:     "3000",
		LogLevel: "debug",
	}
}

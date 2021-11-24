package server

import (
	"log"
	"strings"

	"github.com/ellywynn/http-server/internal/app/store"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	LogLevel string
	Store    *store.Config
}

func NewConfig(configPath string) *Config {
	pathParts := strings.Split(configPath, "/")
	filepath := strings.Join(pathParts[:len(pathParts)-1], "/")
	filename := pathParts[len(pathParts)-1]

	viper.AddConfigPath(filepath)
	viper.SetConfigName(filename)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading server config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error while loading .env file: %s", err.Error())
	}

	return &Config{
		Port:     viper.GetString("server.port"),
		LogLevel: viper.GetString("server.log_level"),
		Store:    store.NewConfig(),
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		Port:     "3000",
		LogLevel: "debug",
	}
}

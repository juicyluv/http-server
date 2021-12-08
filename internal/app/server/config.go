package server

import (
	"log"
	"strings"

	"github.com/ellywynn/http-server/internal/app/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port           string
	LogLevel       string
	WriteTimeout   int
	ReadTimeout    int
	MaxHeaderBytes int
	Repository     *repository.Config
}

// Returns new server config
func NewConfig(configPath string) *Config {
	// Select config path and config filename from path
	pathParts := strings.Split(configPath, "/")
	filepath := strings.Join(pathParts[:len(pathParts)-1], "/")
	filename := pathParts[len(pathParts)-1]

	// Set viper configuration and read config
	viper.AddConfigPath(filepath)
	viper.SetConfigName(filename)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading server config: %s", err.Error())
	}

	// Load .env file
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error while loading .env file: %s", err.Error())
	}

	return &Config{
		Port:           viper.GetString("server.port"),
		LogLevel:       viper.GetString("server.log_level"),
		WriteTimeout:   viper.GetInt("server.writeTimeout"),
		ReadTimeout:    viper.GetInt("server.readTimeout"),
		MaxHeaderBytes: viper.GetInt("server.maxHeaderBytes") << 20, // MB
		Repository:     repository.NewConfig(),
	}
}

// Returns default config parameters
func NewDefaultConfig() *Config {
	return &Config{
		Port:           "3000",
		LogLevel:       "debug",
		WriteTimeout:   10,
		ReadTimeout:    10,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
}

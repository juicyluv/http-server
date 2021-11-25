package repository

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DbURL string
}

type DbConfig struct {
	username string
	password string
	dbname   string
	host     string
	port     string
	sslmode  string
}

// Returns config instance
func NewConfig() *Config {
	dbConfig := configureDbConfig()
	return &Config{
		DbURL: getDbURL(dbConfig),
	}
}

// Returns config with fields from config file
func configureDbConfig() *DbConfig {
	return &DbConfig{
		username: viper.GetString("database.username"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   viper.GetString("database.dbname"),
		host:     viper.GetString("database.host"),
		port:     viper.GetString("database.port"),
		sslmode:  viper.GetString("database.sslmode"),
	}
}

// Returns database URL as string with appropriate config fields
func getDbURL(cfg *DbConfig) string {
	return "postgres://" + cfg.username + ":" + cfg.password + "@" +
		cfg.host + ":" + cfg.port + "/" + cfg.dbname + "?sslmode=" + cfg.sslmode
}

package store

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

func NewConfig() *Config {
	dbConfig := configureDbConfig()
	return &Config{
		DbURL: getDbURL(dbConfig),
	}
}

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

func getDbURL(cfg *DbConfig) string {
	return "postgres://" + cfg.username + ":" + cfg.password + "@" +
		cfg.host + ":" + cfg.port + "/" + cfg.dbname + "?sslmode=" + cfg.sslmode
}

package main

import (
	"flag"
	"log"

	"github.com/ellywynn/http-server/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server", "server config file path")

	flag.Parse()
}

// @title Travels API
// @version 1.0
// @description API Server for Travels Web Application

// @host localhost:3000
// @BasePath /api/v1

func main() {
	config := server.NewConfig(configPath)

	s := server.NewServer(config)
	if err := s.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

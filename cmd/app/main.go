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

func main() {
	config := server.NewConfig(configPath)

	s := server.NewServer(config)
	if err := s.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

package server

import (
	"net/http"
	"time"

	v1 "github.com/ellywynn/http-server/internal/app/handler/http/v1"
	"github.com/ellywynn/http-server/internal/app/repository"
	"github.com/ellywynn/http-server/internal/app/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config     *Config
	logger     *logrus.Logger
	router     *mux.Router
	httpServer *http.Server
	repo       *repository.Repository
	service    *service.Service
}

// Create server instance with appropriate config
func NewServer(config *Config) *Server {
	return &Server{
		config:     config,
		logger:     logrus.New(),
		router:     mux.NewRouter(),
		httpServer: configureHttpServer(config),
	}
}

// Run the server
func (s *Server) Run() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configure(); err != nil {
		return err
	}

	s.logger.Info("Starting server on port " + s.config.Port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) configure() error {
	if err := s.configureRepository(); err != nil {
		return err
	}

	s.configureService()
	s.configureRouter()

	return nil
}

// Configure logger level
func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureService() {
	s.service = service.NewService(s.repo)
}

func (s *Server) configureRouter() {
	s.httpServer.Handler = v1.NewHandler(s.service).InitRoutes()
}

func (s *Server) configureRepository() error {
	repo := repository.NewRepository(s.config.Repository)
	if err := repo.Open(); err != nil {
		return err
	}

	s.repo = repo
	return nil
}

func configureHttpServer(config *Config) *http.Server {
	return &http.Server{
		Addr:           ":" + config.Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
}

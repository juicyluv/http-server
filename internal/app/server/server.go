package server

import (
	"io"
	"net/http"
	"time"

	"github.com/ellywynn/http-server/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config     *Config
	logger     *logrus.Logger
	router     *mux.Router
	httpServer *http.Server
	store      *store.Store
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

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting api server")

	return s.httpServer.ListenAndServe()
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

func (s *Server) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Server) configureStore() error {
	st := store.NewStore(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *Server) handleHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "hello")
	}
}

func configureHttpServer(config *Config) *http.Server {
	return &http.Server{
		Addr:           ":" + config.Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
}

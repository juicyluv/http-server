package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sqlx.DB
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sqlx.Open("postgres", s.config.DbURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	logrus.Info("Connected to the database")

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

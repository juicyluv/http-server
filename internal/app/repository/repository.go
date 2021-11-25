package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type Repository struct {
	UserRepository *UserRepository
	config         *Config
	db             *sqlx.DB
}

func NewRepository(config *Config) *Repository {
	return &Repository{
		config: config,
	}
}

func (r *Repository) Open() error {
	db, err := sqlx.Open("postgres", r.config.DbURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	logrus.Info("Connected to the database")

	r.db = db
	r.initRepositories()
	return nil
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) initRepositories() {
	r.UserRepository = NewUserRepository(r.db)
}

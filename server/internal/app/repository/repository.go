package repository

import (
	"github.com/ellywynn/http-server/server/internal/app/models/interfaces"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type Repository struct {
	// Implemented interfaces
	User     interfaces.UserRepository
	Auth     interfaces.AuthRepository
	Travel   interfaces.TravelRepository
	Place    interfaces.PlaceRepository
	UserRole interfaces.UserRoleRepository

	config *Config
	Db     *sqlx.DB
}

// Creates new repository instance with appropriate config
func NewRepository(config *Config) *Repository {
	return &Repository{
		config: config,
	}
}

// Connects to the database
func (r *Repository) Open() error {
	db, err := sqlx.Open("mysql", r.config.DbURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	logrus.Info("Connected to the database")

	r.Db = db
	r.initRepositories()
	return nil
}

// Closes database connection
func (r *Repository) Close() {
	r.Db.Close()
}

func (r *Repository) initRepositories() {
	r.User = NewUserRepository(r.Db)
	r.Auth = NewAuthRepository(r.Db, &r.User)
	r.Travel = NewTravelRepository(r.Db)
	r.Place = NewPlaceRepository(r.Db)
	r.UserRole = NewUserRoleRepository(r.Db)
}

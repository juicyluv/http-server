package repository

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type PlaceRepository struct {
	db *sqlx.DB
}

func NewPlaceRepository(db *sqlx.DB) *PlaceRepository {
	return &PlaceRepository{
		db: db,
	}
}

func (pr *PlaceRepository) Create(place *models.Place) (uint, error) {
	return 0, nil
}

func (pr *PlaceRepository) FindAll() (*[]models.Place, error) {
	return nil, nil
}

func (pr *PlaceRepository) FindById(placeId int) (*models.Place, error) {
	return nil, nil
}

func (pr *PlaceRepository) Update(placeId int, place *models.UpdatePlaceInput) error {
	return nil
}

func (pr *PlaceRepository) Delete(placeId int) error {
	return nil
}

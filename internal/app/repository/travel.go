package repository

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type TravelRepository struct {
	db *sqlx.DB
}

func NewTravelRepository(db *sqlx.DB) *TravelRepository {
	return &TravelRepository{
		db: db,
	}
}

func (tr *TravelRepository) Create(travel *models.Travel) (int, error) {
	return 0, nil
}

func (tr *TravelRepository) FindAll() (*[]models.Travel, error) {
	return nil, nil
}

func (tr *TravelRepository) FindById(travelId int) (*models.Travel, error) {
	return nil, nil
}

func (tr *TravelRepository) Update(travel *models.Travel) error {
	return nil
}

func (tr *TravelRepository) Delete(travelId int) error {
	return nil
}

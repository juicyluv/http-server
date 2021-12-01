package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
)

type PlaceService struct {
	repository interfaces.PlaceRepository
}

func NewPlaceService(repo *interfaces.PlaceRepository) interfaces.PlaceService {
	return &PlaceService{
		repository: *repo,
	}
}

func (ps *PlaceService) Create(place *models.Place) (uint, error) {
	return 0, nil
}

func (ps *PlaceService) GetAll() (*[]models.Place, error) {
	return nil, nil
}

func (ps *PlaceService) GetById(placeId int) (*models.Place, error) {
	return nil, nil
}

func (ps *PlaceService) Update(placeId int, place *models.UpdatePlaceInput) error {
	return nil
}

func (pr *PlaceService) Delete(placeId int) error {
	return nil
}

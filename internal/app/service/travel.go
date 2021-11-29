package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
)

type TravelService struct {
	repository interfaces.TravelRepository
}

func NewTravelService(repo *interfaces.TravelRepository) interfaces.TravelService {
	return &TravelService{
		repository: *repo,
	}
}

func (ts *TravelService) Create(travel *models.Travel) (uint, error) {
	return ts.repository.Create(travel)
}

func (ts *TravelService) GetById(travelId int) (*models.Travel, error) {
	return ts.repository.FindById(travelId)
}

func (ts *TravelService) GetAll() (*[]models.Travel, error) {
	return ts.repository.FindAll()
}

func (ts *TravelService) Update(travel *models.Travel) error {
	return ts.repository.Update(travel)
}

func (ts *TravelService) Delete(travelId int) error {
	return ts.repository.Delete(travelId)
}

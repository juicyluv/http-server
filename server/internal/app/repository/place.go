package repository

import (
	"fmt"
	"strings"

	"github.com/ellywynn/http-server/server/internal/app/models"
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
	var placeId uint
	query := "INSER INTO palaces (name) VALUES ($1) RETURNING id"
	if err := pr.db.QueryRow(query, place.Name).Scan(&placeId); err != nil {
		return 0, err
	}

	return placeId, nil
}

func (pr *PlaceRepository) FindAll() (*[]models.Place, error) {
	var places []models.Place
	query := "SELECT * FROM places"
	if err := pr.db.Select(&places, query); err != nil {
		return nil, err
	}

	return &places, nil
}

func (pr *PlaceRepository) FindById(placeId int) (*models.Place, error) {
	var place models.Place
	query := "SELECT * FROM places WHERE id = $1"
	if err := pr.db.Get(&place, query, placeId); err != nil {
		return nil, err
	}

	return &place, nil
}

func (pr *PlaceRepository) Update(placeId int, place *models.UpdatePlaceInput) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if place.Name != nil {
		values = append(values, fmt.Sprintf("name=%d", argId))
		args = append(args, place.Name)
		argId++
	}

	valuesQuery := strings.Join(values, ", ")
	args = append(args, placeId)

	query := fmt.Sprintf("UPDATE places SET %s WHERE id=$%d", valuesQuery, argId)
	_, err := pr.db.Exec(query, args...)

	return err
}

func (pr *PlaceRepository) Delete(placeId int) error {
	query := "DELETE FROM places WHERE id = $1"
	_, err := pr.db.Exec(query, placeId)

	return err
}

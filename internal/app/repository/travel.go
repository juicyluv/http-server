package repository

import (
	"fmt"
	"strings"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
	"github.com/jmoiron/sqlx"
)

type TravelRepository struct {
	db  *sqlx.DB
	cld interfaces.CloudinaryService
}

func NewTravelRepository(db *sqlx.DB, cld interfaces.CloudinaryService) *TravelRepository {
	return &TravelRepository{
		db:  db,
		cld: cld,
	}
}

func (tr *TravelRepository) Create(travel *models.Travel) (uint, error) {
	var travelId uint
	query := `INSERT INTO travels
			  VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)
			  RETURNING id`

	err := tr.db.QueryRow(
		query,
		travel.Title,
		travel.DurationDays,
		travel.Price,
		travel.PartySize,
		travel.Complexity,
		travel.Description,
		travel.Date,
		travel.Place,
	).Scan(&travelId)
	if err != nil {
		return 0, err
	}

	return travelId, nil
}

func (tr *TravelRepository) FindAll() (*[]models.Travel, error) {
	var travels []models.Travel
	query := "SELECT t.*, p.name as place FROM travels t INNER JOIN places p ON t.place = p.id"

	if err := tr.db.Select(&travels, query); err != nil {
		return nil, err
	}

	for i := range travels {
		travels[i].FormatDate()
	}

	return &travels, nil
}

func (tr *TravelRepository) FindById(travelId int) (*models.Travel, error) {
	var travel models.Travel
	query := `
		SELECT t.*, p.name as place FROM travels t
		INNER JOIN places p ON t.place = p.id
		WHERE t.id = $1
	`
	err := tr.db.Get(&travel, query, travelId)
	travel.FormatDate()

	if err != nil {
		return nil, err
	}

	return &travel, nil
}

func (tr *TravelRepository) Update(travelId int, travel *models.UpdateTravelInput) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if travel.Title != nil {
		values = append(values, fmt.Sprintf("title=$%d", argId))
		args = append(args, *travel.Title)
		argId++
	}

	if travel.DurationDays != nil {
		values = append(values, fmt.Sprintf("duration_days=$%d", argId))
		args = append(args, *travel.DurationDays)
		argId++
	}

	if travel.Price != nil {
		values = append(values, fmt.Sprintf("price=$%d", argId))
		args = append(args, *travel.Price)
		argId++
	}

	if travel.PartySize != nil {
		values = append(values, fmt.Sprintf("party_size=$%d", argId))
		args = append(args, *travel.PartySize)
		argId++
	}

	if travel.Complexity != nil {
		values = append(values, fmt.Sprintf("complexity=$%d", argId))
		args = append(args, *travel.Complexity)
		argId++
	}

	if travel.Description != nil {
		values = append(values, fmt.Sprintf("description=$%d", argId))
		args = append(args, *travel.Description)
		argId++
	}

	if travel.Date != nil {
		values = append(values, fmt.Sprintf("date=$%d", argId))
		args = append(args, *travel.Date)
		argId++
	}

	if travel.Place != nil {
		values = append(values, fmt.Sprintf("place=$%d", argId))
		args = append(args, *travel.Place)
		argId++
	}

	if travel.ImageURL != nil {
		values = append(values, fmt.Sprintf("image_url=$%d", argId))
		args = append(args, *travel.ImageURL)
		argId++
	}

	if travel.Image != nil {
		values = append(values, fmt.Sprintf("place=$%d", argId))
		args = append(args, *travel.Place)
		argId++
	}

	valuesQuery := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE travels SET %s WHERE id = $%d", valuesQuery, argId)
	args = append(args, travelId)

	_, err := tr.db.Exec(query, args...)
	return err
}

func (tr *TravelRepository) Delete(travelId int) error {
	query := "DELETE FROM travels WHERE id = $1"
	_, err := tr.db.Exec(query, travelId)
	return err
}

package repository

import (
	"database/sql"

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
	query := "SELECT * FROM travels"
	rows, err := tr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var travel models.Travel
	for rows.Next() {
		rows.Scan(
			&travel.Id,
			&travel.Title,
			&travel.DurationDays,
			&travel.Price,
			&travel.PartySize,
			&travel.Complexity,
			&travel.Description,
			&travel.Date,
			&travel.Place,
		)
		travels = append(travels, travel)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &travels, nil
}

func (tr *TravelRepository) FindById(travelId int) (*models.Travel, error) {
	var travel models.Travel
	query := "SELECT * FROM travels WHERE id = $1"
	err := tr.db.QueryRow(query, travelId).Scan(
		&travel.Id,
		&travel.Title,
		&travel.DurationDays,
		&travel.Price,
		&travel.PartySize,
		&travel.Complexity,
		&travel.Description,
		&travel.Date,
		&travel.Place,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &travel, nil
}

func (tr *TravelRepository) Update(travel *models.Travel) error {
	return nil
}

func (tr *TravelRepository) Delete(travelId int) error {
	query := "DELETE FROM travels WHERE id = $1"
	_, err := tr.db.Exec(query, travelId)
	return err
}

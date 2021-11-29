package repository_test

import (
	"testing"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/repository"
	"github.com/lane-c-wagner/go-tinydate"
	"github.com/stretchr/testify/assert"
)

func TestTravelRepository_Create(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := &models.Travel{
		Title:        "Title",
		DurationDays: 10,
		Price:        200,
		PartySize:    10,
		Complexity:   4,
		Place:        1,
		Description:  "desc",
		Date:         tinydate.Now().ToTime(),
	}

	travelId, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId)
}

func TestTravelRepository_FindAll(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := &models.Travel{
		Title:        "Title",
		DurationDays: 10,
		Price:        200,
		PartySize:    10,
		Complexity:   4,
		Place:        1,
		Description:  "desc",
		Date:         tinydate.Now().ToTime(),
	}

	travelId1, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId1)

	travel.Title = "Title 2"
	travel.Price = 5000

	travelId2, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId2)

	travels, err := r.Travel.FindAll()

	var wantedType *[]models.Travel

	assert.NoError(t, err)
	assert.IsType(t, wantedType, travels)
	assert.Equal(t, len(*travels), 2)
}

func TestTravelRepository_FindById(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := &models.Travel{
		Title:        "Title",
		DurationDays: 10,
		Price:        200,
		PartySize:    10,
		Complexity:   4,
		Place:        1,
		Description:  "desc",
		Date:         tinydate.Now().ToTime(),
	}

	travelId, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId)

	travel.Id = uint(travelId)

	tr, err := r.Travel.FindById(int(travelId))
	assert.NoError(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, tr.Id, travelId)
}

func TestTravelRepository_Delete(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := &models.Travel{
		Title:        "Title",
		DurationDays: 10,
		Price:        200,
		PartySize:    10,
		Complexity:   4,
		Place:        1,
		Description:  "desc",
		Date:         tinydate.Now().ToTime(),
	}

	travelId, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId)

	err = r.Travel.Delete(int(travelId))
	assert.NoError(t, err)

	tr, err := r.Travel.FindById(int(travelId))
	assert.NoError(t, err)
	assert.Nil(t, tr)
}

package repository_test

import (
	"testing"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/ellywynn/http-server/server/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestTravelRepository_Create(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := models.TestTravel(t)

	travelId, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId)
}

func TestTravelRepository_FindAll(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := models.TestTravel(t)

	travelId1, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId1)

	travel.Title = "Title 2"
	travel.Price = 5000

	travelId2, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId2)

	travels, err := r.Travel.FindAll(100, 1, 0, "title")

	var wantedType *[]models.Travel

	assert.NoError(t, err)
	assert.IsType(t, wantedType, travels)
	assert.Equal(t, len(*travels), 2)
}

func TestTravelRepository_FindById(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := models.TestTravel(t)

	travelId, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId)

	travel.Id = uint(travelId)

	tr, err := r.Travel.FindById(int(travelId))
	assert.NoError(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, tr.Id, travelId)
}

func TestTravelRepository_Update(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := models.TestTravel(t)

	travelId, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId)

	title := "hello world"
	days := 5
	partySize := 5
	complexity := 3
	description := "new description"

	travelToUpdate := &models.UpdateTravelInput{
		Title:        &title,
		DurationDays: &days,
		PartySize:    &partySize,
		Complexity:   &complexity,
		Description:  &description,
	}

	err = r.Travel.Update(int(travelId), travelToUpdate)
	assert.NoError(t, err)

	updatedTravel, _ := r.Travel.FindById(int(travelId))

	assert.Equal(t, updatedTravel.Complexity, complexity)
	assert.Equal(t, updatedTravel.Title, title)
	assert.Equal(t, updatedTravel.DurationDays, days)
	assert.Equal(t, updatedTravel.Description, description)
	assert.Equal(t, *updatedTravel.PartySize, partySize)
}

func TestTravelRepository_Delete(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("travels")

	travel := models.TestTravel(t)

	travelId, err := r.Travel.Create(travel)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, travelId)

	err = r.Travel.Delete(int(travelId))
	assert.NoError(t, err)

	tr, err := r.Travel.FindById(int(travelId))
	assert.Error(t, err)
	assert.Nil(t, tr)
}

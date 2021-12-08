package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPlace(c *gin.Context) {
	var input models.Place
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	placeId, err := h.service.Place.Create(&input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]uint{"id": placeId})
}

func (h *Handler) getAllPlaces(c *gin.Context) {
	var places *[]models.Place
	places, err := h.service.Place.GetAll()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &places)
}

func (h *Handler) getPlaceById(c *gin.Context) {
	placeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid place id")
		return
	}

	place, err := h.service.Place.GetById(placeId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if place == nil {
		errorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("place with id %d not found", placeId))
		return
	}

	c.JSON(http.StatusOK, place)
}

func (h *Handler) updatePlace(c *gin.Context) {
	var input models.UpdatePlaceInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	placeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid place id")
		return
	}

	if err := h.service.Place.Update(placeId, &input); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) deletePlace(c *gin.Context) {
	placeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid place id")
		return
	}

	if err := h.service.Place.Delete(placeId); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

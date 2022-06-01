package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTravel(c *gin.Context) {
	var input models.Travel

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	travelId, err := h.service.Travel.Create(&input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]uint{"id": travelId})
}

func (h *Handler) getAllTravels(c *gin.Context) {
	var travels *[]models.Travel

	q := c.Request.URL.Query()
	count, _ := strconv.Atoi(q.Get("count"))
	page, _ := strconv.Atoi(q.Get("page"))

	if count == 0 {
		count = 10
	}

	if page == 0 {
		page = 1
	}

	travels, err := h.service.Travel.GetAll(count, page)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &travels)
}

func (h *Handler) getTravelById(c *gin.Context) {
	travelId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid travel id")
		return
	}

	travel, err := h.service.Travel.GetById(travelId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if travel == nil {
		errorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("travel with id %d not found", travelId))
		return
	}

	c.JSON(http.StatusOK, travel)
}

func (h *Handler) updateTravel(c *gin.Context) {
	var input models.UpdateTravelInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	travelId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid travel id")
		return
	}

	if err := h.service.Travel.Update(travelId, &input); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) deleteTravel(c *gin.Context) {
	travelId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid travel id")
		return
	}

	if err := h.service.Travel.Delete(travelId); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) uploadTravelImage(c *gin.Context) {
	filepath, err := h.parseFormFile(c, "travel_image")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	title := c.Request.FormValue("title")
	if title == "" {
		errorResponse(c, http.StatusBadRequest, "title cannot be empty")
		return
	}

	url, err := h.service.Cld.UploadImage(title, filepath, "travels")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{"URL": url})
}

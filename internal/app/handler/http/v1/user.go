package v1

import (
	"net/http"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userId, err := h.service.UserService.Create(&input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]int{"id": userId})
}

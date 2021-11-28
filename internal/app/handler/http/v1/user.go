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

	if err := input.Validate(); err != nil {
		errorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userId, err := h.service.User.SignUp(&input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]int{"id": userId})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	var users *[]models.User
	users, err := h.service.User.GetAll()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"users": &users})
}

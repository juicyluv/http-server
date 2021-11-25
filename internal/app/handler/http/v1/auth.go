package v1

import (
	"net/http"

	"github.com/ellywynn/http-server/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func (h *Handler) login(c *gin.Context) {
	var input repository.LoginStruct
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.service.Auth.Login(&input); err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

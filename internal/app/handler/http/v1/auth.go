package v1

import (
	"net/http"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(c *gin.Context) {
	var input models.AuthLoginStruct
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := h.service.Auth.LogIn(&input)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Set user session
	session := sessions.Default(c)

	session.Set("user_id", user.Id)
	session.Set("email", user.Email)
	session.Set("username", user.Username)

	session.Save()

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) signOut(c *gin.Context) {

}

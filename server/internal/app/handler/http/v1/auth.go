package v1

import (
	"database/sql"
	"net/http"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/ellywynn/http-server/server/internal/app/models/interfaces"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// If creating user with role, require admin role
	if input.Role != nil {
		if !isAdmin(h, c) {
			errorResponse(c, http.StatusForbidden, "you have not access to this action")
			return
		}
	}

	emailTaken, err := h.service.User.GetByEmail(input.Email)
	if err != nil && err != sql.ErrNoRows {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if emailTaken != nil {
		errorResponse(c, http.StatusBadRequest, "email taken")
		return
	}

	usernameTaken, err := h.service.User.GetByUsername(input.Username)
	if err != nil && err != sql.ErrNoRows {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if usernameTaken != nil {
		errorResponse(c, http.StatusBadRequest, "username taken")
		return
	}

	userId, err := h.service.User.SignUp(&input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]int{"id": userId})
}

func (h *Handler) signIn(c *gin.Context) {
	session, err := h.sessionStore.Get(c.Request, coockieName)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// User already authenticated
	if _, exist := session.Values["user_id"]; exist {
		errorResponse(c, http.StatusForbidden, "you are already auhorized")
		return
	}

	var input interfaces.AuthLoginStruct
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := h.service.Auth.LogIn(input)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Set user session
	session.Values["user_id"] = user.Id
	session.Values["email"] = user.Email
	session.Values["username"] = user.Username
	session.Values["role"] = user.Role

	if err := session.Save(c.Request, c.Writer); err != nil {
		errorResponse(c, http.StatusInternalServerError, "unable save user session")
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) signOut(c *gin.Context) {
	session, err := h.sessionStore.Get(c.Request, coockieName)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Clear user session
	session.Options.MaxAge = -1
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "failed to delete session")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

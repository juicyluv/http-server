package v1

import (
	"html/template"
	"net/http"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

const (
	templatesPath = "client/"
)

func (h *Handler) renderIndex(c *gin.Context) {
	tmpl, err := template.ParseFiles(templatesPath + "index.html")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	travels, err := h.service.Travel.GetAll()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := getSessionUserStruct(h, c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tmpl.Execute(c.Writer, struct {
		Travels []models.Travel
		User    *models.User
	}{
		Travels: *travels,
		User:    user,
	})
}

func (h *Handler) renderSignUp(c *gin.Context) {
	tmpl, err := template.ParseFiles(templatesPath + "register.html")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tmpl.Execute(c.Writer, nil)
}

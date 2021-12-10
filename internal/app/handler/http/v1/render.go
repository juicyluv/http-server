package v1

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

var (
	t = template.Must(template.ParseGlob("client/*.html"))
)

func (h *Handler) renderIndex(c *gin.Context) {
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

	var isAuth bool = false
	var username, email, role string
	var userId int
	if user != nil {
		isAuth = true
		username = user.Username
		email = user.Email
		role = *user.Role
		userId = user.Id
	}

	if err = t.ExecuteTemplate(c.Writer, "index.html", struct {
		Travels   []models.Travel
		PageTitle string
		UserId    int
		Username  string
		Email     string
		Role      string
		IsAuth    bool
	}{
		PageTitle: "TRAVELS",
		Travels:   *travels,
		UserId:    userId,
		Username:  username,
		Email:     email,
		Role:      role,
		IsAuth:    isAuth,
	}); err != nil {
		fmt.Println(err.Error())
	}
}

func (h *Handler) renderSignUp(c *gin.Context) {
	t.ExecuteTemplate(c.Writer, "signup.html", nil)
}

func (h *Handler) renderSignIn(c *gin.Context) {
	t.ExecuteTemplate(c.Writer, "signin.html", nil)
}

package v1

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) renderIndex(c *gin.Context) {
	t := template.Must(template.ParseFiles("../client/base.html", "../client/index.html"))

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

func (h *Handler) renderTravel(c *gin.Context) {
	travelId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid travel id")
		return
	}

	t := template.Must(template.ParseFiles("../client/base.html", "../client/travel.html"))

	travel, err := h.service.Travel.GetById(travelId)
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

	if err = t.ExecuteTemplate(c.Writer, "travel.html", struct {
		Travel    models.Travel
		PageTitle string
		UserId    int
		Username  string
		Email     string
		Role      string
		IsAuth    bool
	}{
		PageTitle: travel.Title,
		Travel:    *travel,
		UserId:    userId,
		Username:  username,
		Email:     email,
		Role:      role,
		IsAuth:    isAuth,
	}); err != nil {
		fmt.Println(err.Error())
	}
}

func (h *Handler) renderOrders(c *gin.Context) {
	t := template.Must(template.ParseFiles("../client/base.html", "../client/orders.html"))

	user, err := getSessionUserStruct(h, c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	orders, err := h.service.User.GetTravels(user.Id)
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

	if err = t.ExecuteTemplate(c.Writer, "orders.html", struct {
		Orders    []models.Travel
		PageTitle string
		UserId    int
		Username  string
		Email     string
		Role      string
		IsAuth    bool
	}{
		PageTitle: "Мои заказы",
		Orders:    *orders,
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
	t := template.Must(template.ParseFiles("../client/base.html", "../client/signup.html"))
	if err := t.ExecuteTemplate(c.Writer, "signup.html", nil); err != nil {
		fmt.Println(err.Error())
	}
}

func (h *Handler) renderSignIn(c *gin.Context) {
	t := template.Must(template.ParseFiles("../client/base.html", "../client/signin.html"))
	if err := t.ExecuteTemplate(c.Writer, "signin.html", nil); err != nil {
		fmt.Println(err.Error())
	}
}

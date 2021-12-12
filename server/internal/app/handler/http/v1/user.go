package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	var users *[]models.User
	users, err := h.service.User.GetAll()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &users)
}

func (h *Handler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.service.User.GetById(userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		errorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("user with id %d not found", userId))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	var input models.UserUpdateInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := h.service.User.Update(userId, &input); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := h.service.User.Delete(userId); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) addUserTravel(c *gin.Context) {
	userId, err := getSessionUserId(h, c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	travelId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid travel id")
		return
	}

	if err := h.service.User.AddTravel(userId, travelId); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) addUserTravelByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	travelId, err := strconv.Atoi(c.Param("travel_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid travel id")
		return
	}

	if err := h.service.User.AddTravel(userId, travelId); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) getUserTravels(c *gin.Context) {
	userId, err := getSessionUserId(h, c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	travels, err := h.service.User.GetTravels(userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, travels)
}

func (h *Handler) getUserTravelsByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	travels, err := h.service.User.GetTravels(userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, travels)
}

func (h *Handler) removeUserTravel(c *gin.Context) {
	userId, err := getSessionUserId(h, c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	travelId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid travel id")
		return
	}

	if err := h.service.User.RemoveTravel(userId, travelId); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

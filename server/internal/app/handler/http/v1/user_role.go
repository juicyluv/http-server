package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUserRole(c *gin.Context) {
	var input models.UserRole
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	roleId, err := h.service.UserRole.Create(&input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]uint{"id": roleId})
}

func (h *Handler) getAllUserRoles(c *gin.Context) {
	var roles *[]models.UserRole
	roles, err := h.service.UserRole.GetAll()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &roles)
}

func (h *Handler) getUserRoleById(c *gin.Context) {
	roleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid role id")
		return
	}

	role, err := h.service.UserRole.GetById(roleId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if role == nil {
		errorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("role with id %d not found", roleId))
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *Handler) updateUserRole(c *gin.Context) {
	var input models.UpdateUserRoleInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	roleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid role id")
		return
	}

	if err := h.service.UserRole.Update(roleId, &input); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) deleteUserRole(c *gin.Context) {
	roleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid role id")
		return
	}

	if err := h.service.UserRole.Delete(roleId); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

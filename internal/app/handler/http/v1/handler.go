package v1

import (
	"github.com/ellywynn/http-server/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.POST("/", h.createUser)
		}
	}

	return router
}

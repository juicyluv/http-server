package v1

import (
	"os"

	"github.com/ellywynn/http-server/internal/app/handler/http/middleware"
	"github.com/ellywynn/http-server/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type Handler struct {
	service      *service.Service
	sessionStore sessions.Store
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:      service,
		sessionStore: sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY"))),
	}
}

// Initializes routes and returns handler
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.GET("/logout", h.signOut)
	}

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/", middleware.Authenticate(h.sessionStore), h.getAllUsers)
			users.POST("/", h.signUp)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	return router
}

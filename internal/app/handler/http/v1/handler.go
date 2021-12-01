package v1

import (
	"os"

	"github.com/ellywynn/http-server/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

const (
	coockieName = "travels"
)

type Handler struct {
	service      *service.Service
	sessionStore sessions.Store
}

func NewHandler(service *service.Service) *Handler {
	coockieOptions := sessions.Options{
		HttpOnly: viper.GetBool("sessions.httpOnly"),
		MaxAge:   viper.GetInt("sessions.maxAge") * 60 * 60 * 24, // days
		Secure:   viper.GetBool("sessions.secure"),
		Path:     "/",
	}

	coockieStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	coockieStore.Options = &coockieOptions

	return &Handler{
		service:      service,
		sessionStore: coockieStore,
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
			users.GET("/", authenticate(h.sessionStore), h.getAllUsers)
			users.POST("/", h.signUp)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)

			travels := users.Group("/travels")
			{
				travels.GET("/", h.getUserTravels)
				travels.POST("/:id", h.addUserTravel)
				travels.GET("/:id", h.getTravelById)
				travels.DELETE("/:id", h.removeUserTravel)
			}

			roles := users.Group("/roles")
			{
				roles.GET("/", h.getAllUserRoles)
				roles.POST("/", h.createUserRole)
				roles.GET("/:id", h.getUserRoleById)
				roles.PUT("/:id", h.updateUserRole)
				roles.DELETE("/:id", h.deleteUserRole)
			}
		}

		travels := api.Group("/travels")
		{
			travels.GET("/", h.getAllTravels)
			travels.POST("/", h.createTravel)
			travels.GET("/:id", h.getTravelById)
			travels.PUT("/:id", h.updateTravel)
			travels.DELETE("/:id", h.deleteTravel)
		}

		places := api.Group("/places")
		{
			places.GET("/", h.getAllPlaces)
			places.POST("/", h.createPlace)
			places.GET("/:id", h.getPlaceById)
			places.PUT("/:id", h.updatePlace)
			places.DELETE("/:id", h.deletePlace)
		}
	}

	return router
}

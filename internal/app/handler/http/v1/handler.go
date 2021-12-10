package v1

import (
	"os"

	_ "github.com/ellywynn/http-server/docs"
	"github.com/ellywynn/http-server/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.GET("/logout", h.signOut)
	}

	// API routes
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", authenticate(h.sessionStore), h.getAllUsers)
			users.POST("", h.signUp)
			users.GET("/:id", h.getUserById)
			users.GET("/:id/travels", h.getUserTravelsByUserId)
			users.POST("/:user_id/travels/:travel_id", h.addUserTravelByUserId)
			users.PUT("/:id", h.updateUser)
			users.DELETE(
				"/:id",
				authenticate(h.sessionStore),
				requireRole("Admin", h.sessionStore),
				h.deleteUser,
			)

			travels := users.Group("/travels")
			{
				travels.POST("/:id", h.addUserTravel)
				travels.GET("", h.getUserTravels)
				travels.DELETE("/:id", h.removeUserTravel)
			}

			roles := users.Group("/roles")
			{
				roles.GET("", h.getAllUserRoles)
				roles.POST("", h.createUserRole)
				roles.GET("/:id", h.getUserRoleById)
				roles.PUT("/:id", h.updateUserRole)
				roles.DELETE("/:id", h.deleteUserRole)
			}
		}

		travels := api.Group("/travels")
		{
			travels.GET("", h.getAllTravels)
			travels.POST("", h.createTravel)
			travels.GET("/:id", h.getTravelById)
			travels.PUT("/:id", h.updateTravel)
			travels.DELETE("/:id", h.deleteTravel)
		}

		places := api.Group("/places")
		{
			places.GET("", h.getAllPlaces)
			places.POST("", h.createPlace)
			places.GET("/:id", h.getPlaceById)
			places.PUT("/:id", h.updatePlace)
			places.DELETE("/:id", h.deletePlace)
		}
	}

	// Swagger documentation page
	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	index := router.Group("/")
	{
		index.GET("", h.renderIndex)
		// root.GET("sign-in", h.renderSignIn)
		index.GET("sign-up", h.renderSignUp)
		// root.GET("orders", h.renderOrders)
		// root.GET("travels/:id", h.renderTravel)
	}

	// Static files
	router.Static("/static", "./client/static")

	return router
}

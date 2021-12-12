package v1

import (
	"os"

	_ "github.com/ellywynn/http-server/server/docs"
	"github.com/ellywynn/http-server/server/internal/app/service"
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
		auth.POST(
			"/sign-in",
			alreadyAuthenticated(h.sessionStore),
			h.signIn)

		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-out", h.signOut)
	}

	// API routes
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET(
				"",
				authenticate(h.sessionStore),
				h.getAllUsers)

			users.POST("", h.signUp)
			users.GET("/:id", h.getUserById)
			users.GET("/:id/travels", h.getUserTravelsByUserId)

			users.POST(
				"/:user_id/travels/:travel_id",
				authenticate(h.sessionStore),
				h.addUserTravelByUserId)

			users.PUT(
				"/:id",
				authenticate(h.sessionStore),
				requireRole("Admin", h.sessionStore),
				h.updateUser)

			users.DELETE(
				"/:id",
				authenticate(h.sessionStore),
				requireRole("Admin", h.sessionStore),
				h.deleteUser,
			)

			travels := users.Group("/travels")
			{
				travels.POST(
					"/:id",
					authenticate(h.sessionStore),
					h.addUserTravel)

				travels.GET("", h.getUserTravels)

				travels.DELETE(
					"/:id",
					authenticate(h.sessionStore),
					h.removeUserTravel)
			}

			roles := users.Group("/roles")
			{
				roles.GET("", h.getAllUserRoles)

				roles.POST(
					"",
					authenticate(h.sessionStore),
					requireRole("Admin", h.sessionStore),
					h.createUserRole)
				roles.GET("/:id", h.getUserRoleById)

				roles.PUT("/:id",
					authenticate(h.sessionStore),
					requireRole("Admin", h.sessionStore),
					h.updateUserRole)

				roles.DELETE("/:id",
					authenticate(h.sessionStore),
					requireRole("Admin", h.sessionStore),
					h.deleteUserRole)
			}
		}

		travels := api.Group("/travels")
		{
			travels.GET("", h.getAllTravels)
			travels.POST(
				"",
				authenticate(h.sessionStore),
				requireRole("Admin",
					h.sessionStore),
				h.createTravel)

			travels.GET("/:id", h.getTravelById)

			travels.PUT(
				"/:id",
				authenticate(h.sessionStore),
				requireRole("Admin", h.sessionStore),
				h.updateTravel)

			travels.DELETE(
				"/:id",
				authenticate(h.sessionStore),
				requireRole("Admin", h.sessionStore),
				h.deleteTravel)
		}

		places := api.Group("/places")
		{
			places.GET("", h.getAllPlaces)

			places.POST(
				"",
				authenticate(h.sessionStore),
				requireRole("Admin",
					h.sessionStore),
				h.createPlace)

			places.GET("/:id", h.getPlaceById)

			places.PUT(
				"/:id",
				authenticate(h.sessionStore),
				requireRole("Admin",
					h.sessionStore),
				h.updatePlace)

			places.DELETE(
				"/:id",
				authenticate(h.sessionStore),
				requireRole("Admin",
					h.sessionStore),
				h.deletePlace)
		}
	}

	// Swagger documentation page
	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	index := router.Group("/")
	{
		index.GET("", h.renderIndex)
		index.GET(
			"sign-in",
			alreadyAuthenticated(h.sessionStore),
			h.renderSignIn)

		index.GET(
			"sign-up",
			alreadyAuthenticated(h.sessionStore),
			h.renderSignUp)

		index.GET("sign-out", h.signOut)
		index.GET("travel/:id", h.renderTravel)

		index.GET(
			"orders",
			authenticate(h.sessionStore),
			h.renderOrders)
	}

	// Static files
	router.Static("/static", "../client/static")

	return router
}

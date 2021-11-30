package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func Authenticate(sessionStore sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := sessionStore.Get(c.Request, "cookie-name")
		if _, exists := session.Values["user_id"]; !exists {
			c.JSON(http.StatusForbidden, map[string]string{
				"error": "you need to authorize to access this resource",
			})
			c.Abort()
		}
	}
}

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func authenticate(sessionStore sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := sessionStore.Get(c.Request, coockieName)
		if _, exists := session.Values["user_id"]; !exists {
			c.JSON(http.StatusForbidden, map[string]string{
				"error": "you need to authorize to access this resource",
			})
			c.Abort()
		}
	}
}

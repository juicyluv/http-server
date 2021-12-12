package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// Require authentication
func authenticate(sessionStore sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := sessionStore.Get(c.Request, coockieName)
		if _, exists := session.Values["user_id"]; !exists {
			errorResponse(c, http.StatusForbidden,
				"you need to authorize to access this resource")
			c.Abort()
		}
	}
}

// Require role
func requireRole(role string, sessionStore sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := sessionStore.Get(c.Request, coockieName)
		if session.Values["role"] != role {
			errorResponse(c, http.StatusForbidden, "you have no access to this resource")
			c.Abort()
		}
	}
}

// If user already authenticated, redirect to home page
func alreadyAuthenticated(sessionStore sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := sessionStore.Get(c.Request, coockieName)
		if _, exists := session.Values["user_id"]; exists {
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}
	}
}

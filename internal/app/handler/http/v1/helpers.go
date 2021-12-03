package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func getSessionUserId(h *Handler, c *gin.Context) (int, error) {
	session, _ := h.sessionStore.Get(c.Request, coockieName)
	userId, exists := session.Values["user_id"]
	if !exists {
		return 0, errors.New("you need to sign in")
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return 0, errors.New("user id should be an integer")
	}

	return userIdInt, nil
}

func isAdmin(h *Handler, c *gin.Context) bool {
	session, _ := h.sessionStore.Get(c.Request, coockieName)
	return session.Values["role"] == "Admin"
}

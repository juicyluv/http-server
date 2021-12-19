package v1

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/ellywynn/http-server/server/internal/app/models"
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

func getSessionUserStruct(h *Handler, c *gin.Context) (*models.User, error) {
	session, _ := h.sessionStore.Get(c.Request, coockieName)
	userId, exists := session.Values["user_id"]
	if !exists {
		return nil, nil
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return nil, errors.New("user id should be an integer")
	}

	username := session.Values["username"]
	email := session.Values["email"]
	role := session.Values["role"]

	usernameString, ok := username.(string)
	if !ok {
		return nil, errors.New("username should be a string")
	}
	emailString, ok := email.(string)
	if !ok {
		return nil, errors.New("email should be a string")
	}
	roleString, ok := role.(string)
	if !ok {
		return nil, errors.New("role should be a string")
	}

	user := &models.User{
		Id:       userIdInt,
		Username: usernameString,
		Email:    emailString,
		Role:     &roleString,
	}

	return user, nil
}

func (h *Handler) parseFormFile(c *gin.Context, formName string) (string, error) {
	file, header, err := c.Request.FormFile(formName)
	if err != nil {
		return "", err
	}
	filename := header.Filename
	filepath := fmt.Sprintf("./tmp/%s", filename)
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

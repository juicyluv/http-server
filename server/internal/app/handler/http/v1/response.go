package v1

import "github.com/gin-gonic/gin"

func errorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, map[string]string{"error": msg})
}

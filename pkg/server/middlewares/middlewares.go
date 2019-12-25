package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Test dummy
func Test(c *gin.Context) {
	c.String(http.StatusOK, "%+v", "backend")
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainPage dummy
func MainPage(c *gin.Context) {
	c.String(http.StatusOK, "%+v", "backend")
}

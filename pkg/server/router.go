package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/TwoFlower3/mng-template/pkg/handlers"
	"github.com/TwoFlower3/mng-template/pkg/logger"
)

// RouterOptions struct for router options
type RouterOptions struct {
	log *logger.Logger
}

// CreateRouter create router
func createRouter(ro RouterOptions) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	{
		log := ro.log
		e.Use(Logger(log))
		e.Use(Recovery(log))
		e.Use(requestID())
		initCORS(e)
		initRoutes(e)
	}
	return e
}

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := uuid.NewV4()
		c.Set("Request-ID", id.String())
		c.Next()
	}
}

func initCORS(e *gin.Engine) {
	cfg := cors.DefaultConfig()
	{
		cfg.AllowAllOrigins = true
		cfg.AllowMethods = append(cfg.AllowMethods, "OPTIONS", "GET")
	}

	e.Use(cors.New(cfg))
}

func initRoutes(e gin.IRouter) {
	e.GET("/", handlers.MainPage)
}

package server

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/TwoFlower3/mng-template/pkg/logger"
)

var timeFormat = "02-01-2006 15:04:05"

// Logger setup logger handler
func Logger(log *logger.Logger) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	gin.Logger()

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		requestTime := start.Format(timeFormat)

		c.Next()

		stop := time.Now()
		latency := stop.Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()

		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		reqID := c.GetString("Request-ID")
		method := c.Request.Method

		entry := log.WithFields(logrus.Fields{
			"Hostname":    hostname,
			"Path":        path,
			"Date":        requestTime,
			"Latency":     latency,
			"Code":        statusCode,
			"IP":          clientIP,
			"User-Agent":  clientUserAgent,
			"Referer":     referer,
			"Data-Length": dataLength,
			"Request-ID":  reqID,
			"Method":      method,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := "HTTP Request"
			if statusCode >= 500 {
				entry.Error(msg)
			} else if statusCode >= 400 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}

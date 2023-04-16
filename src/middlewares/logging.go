package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		reqMethod := c.Request.Method
		client := c.ClientIP()
		url := c.Request.RequestURI
		status := c.Writer.Status()
		logrus.WithFields(
			logrus.Fields{
				"Method":      reqMethod,
				"Client IP":   client,
				"Request URL": url,
				"Status":      status,
			}).Info("HTTP Request")
	}
}

package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/KeviinMoralees/crud-products/pkg/metrics"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		path := c.FullPath()
		if path == "" {
			path = "unmatched"
		}

		statusCode := fmt.Sprintf("%d", c.Writer.Status())
		elapsed := time.Since(start).Seconds()

		metrics.HTTPRequestsTotal.WithLabelValues(c.Request.Method, path, statusCode).Inc()
		metrics.HTTPRequestDuration.WithLabelValues(c.Request.Method, path, statusCode).Observe(elapsed)
	}
}

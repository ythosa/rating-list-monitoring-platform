package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricsMiddleware struct {
	opsProcessed *prometheus.CounterVec
	metricsPath  string
}

func NewMetricsMiddleware(metricsPath string) *MetricsMiddleware {
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events"}, []string{"method", "path", "statuscode"},
	)

	return &MetricsMiddleware{
		opsProcessed: opsProcessed,
		metricsPath:  metricsPath,
	}
}

// Metrics middleware to collect metrics from http requests
func (m *MetricsMiddleware) Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		m.opsProcessed.With(prometheus.Labels{
			"method":     c.Request.Method,
			"path":       c.Request.URL.String(),
			"statuscode": strconv.Itoa(c.Writer.Status())},
		).Inc()
	}
}

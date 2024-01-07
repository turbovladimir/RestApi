package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	metrics "github.com/turbovladimir/RestApi/internal/app/metrics"
	"time"
)

func RequestMiddleware(m *metrics.Metrics) gin.HandlerFunc {
	return func(context *gin.Context) {
		tn := time.Now()
		uri := context.Request.RequestURI

		defer func() {
			m.RequestDurationMetric.With(prometheus.Labels{"uri": uri}).Observe(time.Now().Sub(tn).Seconds())
		}()

		m.TotalRequestsMetric.With(prometheus.Labels{"uri": uri}).Inc()

		context.Next()
	}
}

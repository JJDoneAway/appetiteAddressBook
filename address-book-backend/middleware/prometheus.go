package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// @Summary      List all Prometheus metrics
// @Description  Provide a list of all currently known Prometheus metrics
// @ID	         metrics
// @Tags         metrics
// @Produce      plain
// @Success      200  {string}  string "Prometheus metrics line by line"
// @Router       /metrics [get]
func RegisterPrometheus(router *gin.Engine) {
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.0001, 0.0002, 0.0003, 0.0004, 0.0005, 0.0006, 0.0007, 0.0008, 0.0009, 0.001, 0.002, 0.003, 0.004, 0.005, 0.006, 0.007, 0.008, 0.009, 0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09, 0.1, 0.5, 1, 2})

	// set middleware for gin
	m.Use(router)
}

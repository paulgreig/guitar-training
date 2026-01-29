package obs

import (
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	metricsNamespace = "guitar_training"
)

var (
	// Menu selection duration (e.g. time to open scales list or scale detail), by action label.
	menuSelectionDuration *prometheus.HistogramVec
	// View counts by action (scales_list, scale_detail, lessons_list, lesson_detail).
	viewsTotal prometheus.CounterVec
	// Key presses total.
	keyPressesTotal prometheus.Counter
	// Data load results.
	dataLoadTotal prometheus.CounterVec

	reg     *prometheus.Registry
	initProm sync.Once
	server  *http.Server
)

func initPrometheusRegistry() {
	initProm.Do(func() {
		reg = prometheus.NewRegistry()

		// Go runtime: GC, goroutines, memory alloc, etc.
		reg.MustRegister(collectors.NewGoCollector())
		// Process: CPU, RSS, open FDs (where supported).
		reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
			Namespace: metricsNamespace,
		}))

		menuSelectionDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: metricsNamespace,
			Name:      "menu_selection_duration_seconds",
			Help:      "Time taken for a menu selection (e.g. open scales list or scale detail).",
			Buckets:   prometheus.DefBuckets,
		}, []string{"action"})
		reg.MustRegister(menuSelectionDuration)

		viewsTotal = *prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Name:      "views_total",
			Help:      "Total number of views by type (scales_list, scale_detail, lessons_list, lesson_detail).",
		}, []string{"view"})
		reg.MustRegister(viewsTotal)

		keyPressesTotal = prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Name:      "key_presses_total",
			Help:      "Total number of key presses.",
		})
		reg.MustRegister(keyPressesTotal)

		dataLoadTotal = *prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Name:      "data_load_total",
			Help:      "Data load attempts by result (success, error).",
		}, []string{"result"})
		reg.MustRegister(dataLoadTotal)
	})
}

// StartMetricsServer starts an HTTP server that serves /metrics for Prometheus (and Grafana Cloud).
// Port is read from METRICS_PORT (default "9090"). Disabled if METRICS_PORT is empty or "0".
func StartMetricsServer() {
	initPrometheusRegistry()

	port := os.Getenv("METRICS_PORT")
	if port == "" {
		port = "9090"
	}
	if port == "0" {
		Info("metrics server disabled (METRICS_PORT=0)")
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	server = &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		Info("metrics server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Error("metrics server error: %v", err)
		}
	}()
}

// StopMetricsServer closes the metrics HTTP server if it was started.
func StopMetricsServer() {
	if server != nil {
		_ = server.Close()
		server = nil
	}
}

// RecordMenuSelectionDuration records the duration of a menu selection for performance monitoring.
// action should be one of: scales_list, scale_detail, lessons_list, lesson_detail.
func RecordMenuSelectionDuration(action string, duration time.Duration) {
	initPrometheusRegistry()
	menuSelectionDuration.WithLabelValues(action).Observe(duration.Seconds())
}

// IncViewsTotal increments the view counter for Prometheus (call in addition to existing Record*).
func IncViewsTotal(view string) {
	initPrometheusRegistry()
	viewsTotal.WithLabelValues(view).Inc()
}

// IncKeyPressesTotal increments the key-press counter for Prometheus.
func IncKeyPressesTotal() {
	initPrometheusRegistry()
	keyPressesTotal.Inc()
}

// IncDataLoadTotal increments the data load counter with result "success" or "error".
func IncDataLoadTotal(result string) {
	initPrometheusRegistry()
	dataLoadTotal.WithLabelValues(result).Inc()
}

// MetricsPort returns the port number used for the metrics server, or 0 if disabled.
func MetricsPort() int {
	portStr := os.Getenv("METRICS_PORT")
	if portStr == "" {
		portStr = "9090"
	}
	if portStr == "0" {
		return 0
	}
	p, _ := strconv.Atoi(portStr)
	return p
}

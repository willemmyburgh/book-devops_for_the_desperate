package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	connectionsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "telnet_server_connection_total",
		Help: "The total number of connections",
	})
	connectionErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "telnet_server_connection_errors_total",
		Help: "The total number of errors",
	})
	unknownCommands = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telnet_server_unknown_commands_total",
		Help: "The total number of unknown commands entered",
	}, []string{"command"})
)

// MetricServer holds state for our Prometheus metrics server
type MetricServer struct {
	port     string
	endPoint string
	logger   *log.Logger
}

// New creates a new metric server
func New(port string, logger *log.Logger) *MetricServer {
	return &MetricServer{port: port, endPoint: "/metrics", logger: logger}
}

// ListenAndServeMetrics runs our metrics server
func (m *MetricServer) ListenAndServeMetrics() {
	http.Handle(m.endPoint, promhttp.Handler())
	m.logger.Printf("Metrics endpoint listening on %s\n", m.port)
	http.ListenAndServe(m.port, nil)
}

// IncrementConnectionErrors += 1
func (m *MetricServer) IncrementConnectionErrors() {
	connectionErrors.Inc()
}

//IncrementConnectionsProcessed += 1
func (m *MetricServer) IncrementConnectionsProcessed() {
	connectionsProcessed.Inc()
}

// IncrementUnknownCommands += 1
func (m *MetricServer) IncrementUnknownCommands(cmd string) {
	unknownCommands.WithLabelValues(cmd).Inc()
}

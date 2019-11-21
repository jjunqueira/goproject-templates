package metrics

type Metrics interface {
	Render() string
}

// Metrics contains application metrics
type PrometheusMetrics struct{}

// NewPrometheusMetrics constructs a prometheus metrics instance
func NewPrometheusMetrics() (*PrometheusMetrics, error) {
	m := new(Metrics)
	return m, nil
}

// Renders the metrics for consumption by metrics gatherers
func (m *PrometheusMetrics) Render() string {
	return ""
}
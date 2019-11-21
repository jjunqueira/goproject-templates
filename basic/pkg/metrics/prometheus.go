package metrics

type Metrics interface {
	Render() string
}

// Metrics contains application metrics
type PrometheusMetrics struct{}

func NewPrometheusMetrics() (*PrometheusMetrics, error) {
	m := new(Metrics)
	return m, nil
}

func (*PrometheusMetrics m) Render() string {
	return ""
}
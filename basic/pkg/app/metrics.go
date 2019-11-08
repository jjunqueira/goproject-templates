package app

// Metrics contains application metrics
type Metrics struct{}

func newMetrics() (*Metrics, error) {
	m := new(Metrics)
	return m, nil
}

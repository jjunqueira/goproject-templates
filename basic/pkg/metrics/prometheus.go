package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// Create your metrics here
var (
	// MyCounter counts important things
	MyCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "important_things_total",
			Help: "The number of important things that have happened",
		},
	)
)

// Register your metrics here
func init() {
	err := prometheus.Register(MyCounter)
	if err != nil {
		panic(fmt.Sprintf("unable to initialize metrics %v", err))
	}
}

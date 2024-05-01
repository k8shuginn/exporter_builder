package vector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func SetFlags() {

}

type Collector struct {
	promCollectCount *prometheus.CounterVec
}

func NewCollector() *Collector {
	return &Collector{
		promCollectCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "prometheus_collect_count",
				Help: "Number of times the Prometheus collector is called",
			},
			[]string{"method"},
		),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.promCollectCount.With(prometheus.Labels{"method": "Describe"}).Inc()
	c.promCollectCount.Describe(ch)
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.promCollectCount.With(prometheus.Labels{"method": "Collect"}).Inc()
	c.promCollectCount.Collect(ch)
}

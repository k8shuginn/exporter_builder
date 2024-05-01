package instance

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	val1 = "value1"
	val2 = "value2"

	sampleDesc = prometheus.NewDesc(
		"sample_metric",
		"sample metric",
		[]string{"key1", "key2"}, nil)
)

func SetFlags() {
	flag.StringVar(&val1, "val1", val1, "value1")
	flag.StringVar(&val2, "val2", val2, "value2")
}

type Collector struct {
}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- sampleDesc
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(sampleDesc, prometheus.GaugeValue, 1, val1, val2)
}

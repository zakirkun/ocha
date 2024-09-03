package main

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/zakirkun/ocha/utils"
)

var (
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "node_cpu_usage",
			Help: "CPU usage percentage",
		},
		[]string{"cpu"},
	)
	memoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "node_memory_usage",
			Help: "Memory usage percentage",
		},
	)
	loadAverage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "node_load_average",
			Help: "Load average",
		},
	)
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(loadAverage)
}

func LoadAndSetMetrics() {

	percent := utils.CollectMetricsCPU()
	for i, perc := range percent {
		idx := strconv.Itoa(i)
		cpuUsage.WithLabelValues(string(idx)).Set(perc)
	}

	mem := utils.CollectMetricsMemory()
	memoryUsage.Set(mem)

	avg := utils.CollectMetricsAvg()
	loadAverage.Set(avg)
}

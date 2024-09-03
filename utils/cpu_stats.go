package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func CollectMetricsCPU() []float64 {

	procs := make([]float64, 0)
	// CPU Usage
	percentages, err := cpu.Percent(0, true)
	if err == nil {
		procs = append(procs, percentages...)
	}

	return procs
}

func CollectMetricsMemory() float64 {
	// Memory Usage
	v, err := mem.VirtualMemory()
	if err == nil {
		return v.UsedPercent
	}

	return 0
}

func CollectMetricsAvg() float64 {
	// Load Average
	l, err := load.Avg()
	if err == nil {
		return l.Load1
	}

	return 0
}

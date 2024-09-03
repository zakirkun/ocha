package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		for {
			LoadAndSetMetrics()
			time.Sleep(2 * time.Second) // Adjust the interval as needed
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

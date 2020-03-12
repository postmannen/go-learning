package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			upOrDown.Set(1)
			time.Sleep(2 * time.Second)
			upOrDown.Set(0)
		}
	}()
}

var (
	upOrDownOpts = prometheus.GaugeOpts{
		Name: "up_or_down",
		Help: "gives an indication if it is up or down",
	}
	upOrDown = promauto.NewGauge(upOrDownOpts)
)

//var (
//	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "myapp_processed_ops_total",
//		Help: "The total number of processed events",
//	})
//)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

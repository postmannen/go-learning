package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	hdFailures := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)

	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(hdFailures)

	var counter float64

	go func() {
		for {

			hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Set(10 + counter)
			hdFailures.With(prometheus.Labels{"device": "/dev/sdb"}).Set(20 + counter)
			time.Sleep(time.Second * 2)
			counter++
		}
	}()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

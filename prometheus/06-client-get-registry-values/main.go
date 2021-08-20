package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	reg := prometheus.NewRegistry()

	go func() {
		//http.Handle("/metrics", promhttp.Handler())
		//http.ListenAndServe(":2112", nil)
		n, err := net.Listen("tcp", ":28080")
		if err != nil {
			log.Printf("error: failed to open prometheus listen port: %v\n", err)
			os.Exit(1)
		}
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		http.Serve(n, m)
	}()

	apekatt := promauto.With(reg).NewGauge(prometheus.GaugeOpts{
		Name: "apekatt",
		Help: "apekatter",
	})
	apekatt.Set(7.7)

	fmt.Printf("%v\n", reg)

	mfs, err := reg.Gather()
	if err != nil {
		log.Printf("error: gathering: %v\n", mfs)
	}

	for _, mf := range mfs {
		fmt.Printf(" name : %v, type: %T\n", mf.GetName(), mf.GetName())
		m := mf.GetMetric()
		fmt.Printf(" metric : %v, type %T\n", m, m)

		// m2 := prometheus.Gauge(m)
		fmt.Printf(" metric : %v, type %T\n", m[0].Gauge.GetValue(), m[0].Gauge.GetValue())
	}
}

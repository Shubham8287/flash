package main

import (
	"flash/app"
	. "flash/logger"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

const (
	PORT        = ":8080"
	metricsAddr = ":8081"
)

func handleRequests() {
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HomeReq)
	mux.HandleFunc("/isAlive", app.Middleware(app.PingReq))
	mux.HandleFunc("/find", app.FindReq)
	mux.HandleFunc("/insert", app.InsertReq)
	mux.Handle("/metrics", promhttp.Handler())

	serv := std.Handler("", mdlw, mux)
	Log.WithField("PORT", PORT).Info("server_started")
	// Serve our metrics.
	go func() {
		log.Printf("metrics listening at %s", metricsAddr)
		if err := http.ListenAndServe(metricsAddr, promhttp.Handler()); err != nil {
			log.Panicf("error while serving metrics: %s", err)
		}
	}()
	Log.Fatal(http.ListenAndServe(PORT, serv))
}

func main() {
	handleRequests()
}

package main

import (
	"flash/app"
	. "flash/logger"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const PORT = "8080"

func handleRequests() {
	Log.WithField("PORT", PORT).Info("server_started")
	http.HandleFunc("/", app.HomeReq)
	http.HandleFunc("/isAlive", app.Middleware(app.PingReq))
	http.HandleFunc("/find", app.Middleware(app.FindReq))
	http.HandleFunc("/insert", app.InsertReq)
	http.Handle("/metrics", promhttp.Handler())
	Log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func main() {
	// Spin up the server and ready to serve requests.
	handleRequests()
}

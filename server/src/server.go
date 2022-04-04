package main

import (
	"encoding/json"
	"flash/ds"
	. "flash/logger"
	"flash/util"
	"fmt"
	"html"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func homeReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func pingReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Boss is always fine")
}

func findReq(w http.ResponseWriter, r *http.Request) {
	searchBucket := ds.GetDs()
	setupResponse(&w, r)
	prefix := r.FormValue("prefix")
	matches := searchBucket.Find(prefix)
	w.Header().Set("Content-Type", "application/json")
	response := util.Response{
		Prefix:  prefix,
		Matches: matches,
	}
	json.NewEncoder(w).Encode(response)
}

func insertReq(w http.ResponseWriter, r *http.Request) {
	searchBucket := ds.GetDs()
	setupResponse(&w, r)
	prefix := r.FormValue("prefix")
	searchBucket.Insert(prefix)
	w.Header().Set("Content-Type", "application/json")
	response := util.Response{
		Prefix: prefix,
	}
	json.NewEncoder(w).Encode(response)
}
func handleRequests() {
	http.HandleFunc("/", homeReq)
	http.HandleFunc("/isAlive", pingReq)
	http.HandleFunc("/find", findReq)
	http.HandleFunc("/insert", insertReq)
	Log.Info("server_started")
	Log.Info(http.ListenAndServe(":8080", nil))
}

func main() {
	// Spin up the server and ready to serve requests.
	handleRequests()
}

package main

import (
	"encoding/json"
	"flash/ds"
	"flash/util"
	"fmt"
	"html"
	"log"
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

func initStuff() {
	// naming is hard, please update if find better name
	_ = ds.GetDs()
}

func homeReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func pingReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Boss is always fine")
}

func findReq(w http.ResponseWriter, r *http.Request) {
	searchQueries := ds.GetDs()
	setupResponse(&w, r)
	prefix := r.FormValue("prefix")
	matches := searchQueries.Find(prefix)
	w.Header().Set("Content-Type", "application/json")
	response := util.Response{
		Prefix:  prefix,
		Matches: matches,
	}
	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	fmt.Println("Server Started...")
	http.HandleFunc("/", homeReq)
	http.HandleFunc("/isAlive", pingReq)
	http.HandleFunc("/find", findReq)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// Intialize stuff
	initStuff()
	// Spin up the server and ready to serve requests.
	handleRequests()
}

package app

import (
	"encoding/json"
	"flash/ds"
	"flash/util"
	"fmt"
	"html"
	"net/http"
)

func HomeReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func PingReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Boss is always fine")
}

func FindReq(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	searchBucket := ds.GetDs(ctx)
	// Log.WithFields(logrus.Fields{
	// 	"request_id":   reqID,
	// 	"time_elapsed": elapsed_time,
	// }).Info("find_req_received")
	setupResponse(&w, r)
	prefix := r.FormValue("prefix")
	matches := searchBucket.Find(prefix)
	// Log.WithFields(logrus.Fields{
	// 	"request_id":   reqID,
	// 	"time_elapsed": elapsed_time,
	// }).Info("find_matches_received")
	response := util.Response{
		Prefix:  prefix,
		Matches: matches,
	}
	json.NewEncoder(w).Encode(response)
}

func InsertReq(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	searchBucket := ds.GetDs(ctx)
	setupResponse(&w, r)
	prefix := r.FormValue("prefix")
	searchBucket.Insert(prefix)
	response := util.Response{
		Prefix: prefix,
	}
	json.NewEncoder(w).Encode(response)
}

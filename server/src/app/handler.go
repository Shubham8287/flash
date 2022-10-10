package app

import (
	"encoding/json"
	"flash/ds"
	"flash/util"
	"flash/conf"
	"flash/logger"
	"fmt"
	"html"
	"net/http"
	"github.com/sirupsen/logrus"
)

var config conf.Config = conf.Get()

func HomeReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func PingReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Boss is always fine")
}

func FindReq(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	searchBucket := ds.GetDs(ctx)
	setupResponse(&w, r)
	prefix := r.FormValue("prefix")
	matches := searchBucket.Find(prefix)
	if(config.LoggingEnabled == true) {
		logger.Log.WithFields(logrus.Fields{
			"request_id":   util.GetRequestID(ctx),
			"time_elapsed": util.SetElapsedTimeStamp(ctx),
		}).Info("find_matches_received")
	}
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

package app

import (
	"context"
	. "flash/logger"
	"flash/util"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := uuid.New()
		ctx = context.WithValue(ctx, util.REQUEST_ID, reqID.String())
		ctx = context.WithValue(ctx, util.LAST_LOG_TIME, time.Now())
		r = r.WithContext(ctx)
		Log.WithField("request_id", reqID).Info("request_id_generated")
		next.ServeHTTP(w, r)
	})
}

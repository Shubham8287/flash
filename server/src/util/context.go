package util

import (
	"context"
	"time"
)

const REQUEST_ID = "requestID"
const LAST_LOG_TIME = "lastUpdateTimestamp"

func SetElapsedTimeStamp(ctx context.Context) time.Duration {

	last_time := ctx.Value(LAST_LOG_TIME)
	cur_time := time.Now()
	elapsed_time := cur_time.Sub(last_time.(time.Time))
	ctx = context.WithValue(ctx, REQUEST_ID, last_time)

	return elapsed_time
}

func GetRequestID(ctx context.Context) string {

	reqID := ctx.Value(REQUEST_ID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}

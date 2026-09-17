package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int
const (
	requestId ctxKey = iota
)



func RequestId(Next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		// Check if the request already has a request ID
		if requestID == "" {
			// Generate a new request ID
			requestID = uuid.New().String()
		}
		r.Header.Set("X-Request-Id", requestID)
		ctx := context.WithValue(r.Context(), requestId, requestID)
		w.Header().Set("X-Request-Id", requestID)
		Next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func GetRequestId(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestId).(string); ok {
		return requestID
	}
	return ""
}
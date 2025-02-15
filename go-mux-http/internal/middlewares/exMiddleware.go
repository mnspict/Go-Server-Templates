package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

// Logger is a middleware that logs info about incoming requests.
// Various other middlewares like authentication, authorization, etc can be added here.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("At: %v > Method: %v > Path: %v\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
package middleware

import (
	"fmt"
	"net/http"
	"time"
)

var (
	Logger func(next http.Handler) http.Handler
)

func TimerTrace(next http.Handler) http.Handler {
	return Logger(next)
}

func RequestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			defer func() {
				t2 := time.Since(t1)
				fmt.Println(t2)
			}()
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
func init() {
	Logger = RequestLogger()
}

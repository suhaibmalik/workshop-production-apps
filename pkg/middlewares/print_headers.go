package middlewares

import (
	"fmt"
	"net/http"
)

func DumpHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, h := range r.Header {
			for _, v := range h {
				fmt.Printf("HEADER %s: %s\n", k, v)
			}
		}

		next.ServeHTTP(w, r)
	})
}

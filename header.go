package main

import (
	"net/http"
	"strings"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		origin := req.Header.Get("Origin")
		if origin != "" {
			if strings.HasPrefix(origin, "http://") || strings.HasPrefix(origin, "https://") {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Expose-Headers", "Link")
			w.Header().Set("Access-Control-Max-Age", "300") // 5 minutes

			if req.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		next.ServeHTTP(w, req)
	})
}

// Command latency serves a page that measures the round trip to this server.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const addr = ":8080"

type latencyResponse struct {
	Timestamp time.Time `json:"timestamp"`
	UnixMilli int64     `json:"unix_milli"`
}

type healthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

func main() {
	srv := &http.Server{
		Addr:              addr,
		Handler:           logRequests(newMux(time.Now())),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("listening on %s", addr)
	log.Fatal(srv.ListenAndServe())
}

func newMux(started time.Time) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", serveIndex)
	mux.HandleFunc("GET /ping-latency", pingLatency)
	mux.HandleFunc("OPTIONS /ping-latency", pingLatency)
	health := healthHandler(started)
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /healthz", health)
	return mux
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	noCache(w)
	http.ServeFile(w, r, "index.html")
}

// pingLatency returns the server time; the page measures the round trip.
// It allows any origin so the endpoint can be measured from another page.
func pingLatency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	now := time.Now()
	noCache(w)
	writeJSON(w, latencyResponse{Timestamp: now, UnixMilli: now.UnixMilli()})
}

func healthHandler(started time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, healthResponse{
			Status:    "healthy",
			Timestamp: time.Now(),
			Uptime:    time.Since(started).Round(time.Second).String(),
		})
	}
}

func noCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("encode response: %v", err)
	}
}

// logRequests logs page requests; the latency and health endpoints are
// called continuously and are left out.
func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		switch r.URL.Path {
		case "/ping-latency", "/health", "/healthz":
			return
		}
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

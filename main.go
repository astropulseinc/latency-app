package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Response structure for /ping-latency endpoint
type LatencyResponse struct {
	Timestamp      time.Time `json:"timestamp"`
	UnixMilli      int64     `json:"unix_milli"`
	ServerLocation string    `json:"server_location,omitempty"`
}

// Health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

var startTime time.Time

func main() {
	startTime = time.Now()

	// Serve static files (index.html)
	http.HandleFunc("/", serveIndex)

	// Latency measurement endpoint
	http.HandleFunc("/ping-latency", handlePingLatency)

	// Health check endpoint for Kubernetes/monitoring
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/healthz", handleHealth)

	// Start server
	addr := ":8080"
	log.Printf("🚀 Latency App server starting on %s", addr)
	log.Printf("📊 Endpoints:")
	log.Printf("   - http://localhost:8080/          (Web UI)")
	log.Printf("   - http://localhost:8080/ping-latency (API)")
	log.Printf("   - http://localhost:8080/health    (Health Check)")

	if err := http.ListenAndServe(addr, logRequest(http.DefaultServeMux)); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}

// serveIndex serves the main HTML page
func serveIndex(w http.ResponseWriter, r *http.Request) {
	// Only serve index.html for root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Set cache control headers
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	http.ServeFile(w, r, "index.html")
}

// handlePingLatency handles latency measurement requests
func handlePingLatency(w http.ResponseWriter, r *http.Request) {
	// Only accept GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set CORS headers (allow cross-origin requests)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Get current server time with high precision
	currentTime := time.Now()

	// Create response
	response := LatencyResponse{
		Timestamp: currentTime,
		UnixMilli: currentTime.UnixMilli(),
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	// Encode and send response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("⚠️  Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// handleHealth handles health check requests
func handleHealth(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)

	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Uptime:    uptime.Round(time.Second).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// logRequest is a middleware that logs all HTTP requests
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		handler.ServeHTTP(w, r)

		// Skip logging for ping-latency and health check endpoints to avoid log spam
		if r.URL.Path == "/ping-latency" || r.URL.Path == "/health" || r.URL.Path == "/healthz" {
			return
		}

		// Log the request
		duration := time.Since(start)
		log.Printf("%-6s %-20s %s", r.Method, r.URL.Path, duration)
	})
}

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRoutes(t *testing.T) {
	mux := newMux(time.Now())
	tests := []struct {
		method, path string
		want         int
	}{
		{http.MethodGet, "/", http.StatusOK},
		{http.MethodGet, "/ping-latency", http.StatusOK},
		{http.MethodOptions, "/ping-latency", http.StatusNoContent},
		{http.MethodGet, "/health", http.StatusOK},
		{http.MethodGet, "/healthz", http.StatusOK},
		{http.MethodPost, "/ping-latency", http.StatusMethodNotAllowed},
		{http.MethodGet, "/missing", http.StatusNotFound},
	}
	for _, tt := range tests {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))
		if rec.Code != tt.want {
			t.Errorf("%s %s: status %d, want %d", tt.method, tt.path, rec.Code, tt.want)
		}
	}
}

func TestPingLatency(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux(time.Now()).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/ping-latency", nil))
	var got latencyResponse
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.UnixMilli == 0 || got.Timestamp.IsZero() {
		t.Errorf("empty response: %+v", got)
	}
	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("missing CORS header")
	}
}

func TestHealth(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux(time.Now().Add(-time.Minute)).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	var got healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.Status != "healthy" || got.Uptime != "1m0s" {
		t.Errorf("got %+v", got)
	}
}

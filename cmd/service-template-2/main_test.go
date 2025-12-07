package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := HealthResponse{
			Status:    "ok",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Service:   "backend",
		}

		json.NewEncoder(w).Encode(resp)
	})

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		uptime := time.Since(startTime).Round(time.Second)
		resp := ServiceInfo{
			Name:    "backend-service",
			Version: "1.0.0",
			Uptime:  uptime.String(),
		}

		json.NewEncoder(w).Encode(resp)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Welcome to the backend service 🚀"))
	})

	return r
}

func TestHealthEndpoint(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	expectedContentType := "application/json"
	if contentType != expectedContentType {
		t.Errorf("expected content type %s, got %s", expectedContentType, contentType)
	}

	var response HealthResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if response.Status != "ok" {
		t.Errorf("expected status 'ok', got '%s'", response.Status)
	}

	if response.Service != "backend" {
		t.Errorf("expected service 'backend', got '%s'", response.Service)
	}

	if response.Timestamp == "" {
		t.Error("expected timestamp to be set")
	}

	if _, err := time.Parse(time.RFC3339, response.Timestamp); err != nil {
		t.Errorf("timestamp format is invalid: %v", err)
	}
}

func TestInfoEndpoint(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/info", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	expectedContentType := "application/json"
	if contentType != expectedContentType {
		t.Errorf("expected content type %s, got %s", expectedContentType, contentType)
	}

	var response ServiceInfo
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if response.Name != "backend-service" {
		t.Errorf("expected name 'backend-service', got '%s'", response.Name)
	}

	if response.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got '%s'", response.Version)
	}

	if response.Uptime == "" {
		t.Error("expected uptime to be set")
	}
}

func TestRootEndpoint(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedContent := "Welcome to the backend service 🚀"
	if w.Body.String() != expectedContent {
		t.Errorf("expected content '%s', got '%s'", expectedContent, w.Body.String())
	}
}

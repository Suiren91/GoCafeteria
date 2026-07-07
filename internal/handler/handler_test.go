package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/handler"
	"github.com/gin-gonic/gin"
)

func TestHealthCheckOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := handler.NewRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("want: %d, got: %d", http.StatusOK, w.Code)
	}
}

func TestHealthCheckBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := handler.NewRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	var body map[string]string

	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("Unmarshal filed: %v", err)
	}
	if body["message"] != "ok" {
		t.Errorf("want: %v, got: %v", "ok", body["message"])
	}
}

package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if body["message"] != "ok" {
		t.Errorf("want: %v, got: %v", "ok", body["message"])
	}
}

func TestCreateNewMenuOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := handler.NewRouter()

	w := httptest.NewRecorder()
	body := `{"name":"かぼちゃ","description":"ほくほくでうまい","price":150,"stock":100}`

	req := httptest.NewRequest(http.MethodPost, "/menu/new", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("want: %v, got:%v \n body:%s ", http.StatusOK, w.Code, w.Body.String())
	}
}

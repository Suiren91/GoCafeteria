package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/Suiren91/go-cafeteria/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

func TestHealthCheckOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := &MenuServiceMock{}
	r := handler.NewRouter(handler.NewMenuHandler(mock))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("want: %d, got: %d", http.StatusOK, w.Code)
	}
}

func TestHealthCheckBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := &MenuServiceMock{}
	r := handler.NewRouter(handler.NewMenuHandler(mock))

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

func TestCreateNewMenu(t *testing.T) {
	tests := map[string]struct {
		body              string
		createNewMenuFunc func(ctx context.Context, name, description string, price, stock int) (int, error)
		wantStatus        int
		wantBody          string
	}{
		"正常なリクエスト": {
			body: `{"name":"かぼちゃ","description":"ほくほくでうまい","price":150,"stock":100}`,
			createNewMenuFunc: func(ctx context.Context, name string, description string, price int, stock int) (int, error) {
				return 1, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"id":1}`,
		},
		"名前が存在しないリクエスト": {
			body:              `{"description":"ほくほくでうまい","price":150,"stock":100}`,
			createNewMenuFunc: nil,
			wantStatus:        http.StatusBadRequest,
			wantBody:          "",
		},
		"説明が存在しないリクエスト": {
			// 説明は必須でないので存在しなくても成功する
			body: `{"name":"かぼちゃ","price":150,"stock":100}`,
			createNewMenuFunc: func(ctx context.Context, name string, description string, price int, stock int) (int, error) {
				return 2, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"id":2}`,
		},
		"値段が存在しないリクエスト": {
			body:              `{"name":"かぼちゃ","description":"ほくほくでうまい","stock":100}`,
			createNewMenuFunc: nil,
			wantStatus:        http.StatusBadRequest,
			wantBody:          "",
		},
		"在庫が存在しないリクエスト": {
			body:              `{"name":"かぼちゃ","description":"ほくほくでうまい","price":150"}`,
			createNewMenuFunc: nil,
			wantStatus:        http.StatusBadRequest,
			wantBody:          "",
		},
		"名前が空のリクエスト": {
			body:              `{"name":"","description":"ほくほくでうまい","price":150,"stock":100}`,
			createNewMenuFunc: nil,
			wantStatus:        http.StatusBadRequest,
			wantBody:          "",
		},
		"値段が負の値のリクエスト": {
			body: `{"name":"かぼちゃ","description":"ほくほくでうまい","price":-1,"stock":100}`,
			createNewMenuFunc: func(ctx context.Context, name string, description string, price int, stock int) (int, error) {
				return 0, domain.ErrNegativePrice
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"error":"price must be non-negative"}`,
		},
		"在庫が負の値のリクエスト": {
			body: `{"name":"かぼちゃ","description":"ほくほくでうまい","price":150,"stock":-1}`,
			createNewMenuFunc: func(ctx context.Context, name string, description string, price int, stock int) (int, error) {
				return 0, domain.ErrNegativeStock
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"error":"stock must be non-negative"}`,
		},
		"サーバーエラー": {
			body: `{"name":"かぼちゃ","description":"ほくほくでうまい","price":150,"stock":100}`,
			createNewMenuFunc: func(ctx context.Context, name string, description string, price int, stock int) (int, error) {
				return 0, errors.New("CreateNewMenu failure")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"error":"unexpected error"}`,
		},
	}
	gin.SetMode(gin.TestMode)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mock := &MenuServiceMock{CreateNewMenuFunc: tt.createNewMenuFunc}

			r := handler.NewRouter(handler.NewMenuHandler(mock))
			req := httptest.NewRequest(http.MethodPost, "/menu/new", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Fatalf("status: want: %d, got: %d (body: %s)", tt.wantStatus, w.Code, w.Body)
			}
			if tt.wantBody != "" {
				assertJSONEq(t, tt.wantBody, w.Body.String())
			}
		})
	}
}

func assertJSONEq(t *testing.T, want, got string) {
	t.Helper()
	var w, g any
	if err := json.Unmarshal([]byte(want), &w); err != nil {
		t.Fatalf("want: invalid JSON: %v", err)
	}
	if err := json.Unmarshal([]byte(got), &g); err != nil {
		t.Fatalf("got: invalid JSON: %v\nbody: %s", err, got)
	}
	if diff := cmp.Diff(w, g); diff != "" {
		t.Errorf("response JSON mismatch (-want +got):\n%s", diff)
	}
}

package service_test

import (
	"errors"
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/Suiren91/go-cafeteria/internal/service"
)

func TestCreateNewMenu(t *testing.T) {
	tests := map[string]struct {
		id          int
		name        string
		description string
		price       int
		stock       int
		wantErr     error
	}{
		"適切なパラメータでメニューを新規作成できる": {
			id:          1,
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       150,
			stock:       100,
			wantErr:     nil,
		}, "priceが負の数(-1)だとエラー": {
			id:          1,
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       -1,
			stock:       100,
			wantErr:     domain.ErrNegativePrice,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := service.CreateNewMenu(tt.id, tt.name, tt.description, tt.price, tt.stock)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("want: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}

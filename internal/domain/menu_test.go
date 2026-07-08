package domain_test

import (
	"errors"
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/domain"
)

func TestNewMenu(t *testing.T) {
	tests := map[string]struct {
		id          int
		name        string
		description string
		price       int
		stock       int
		wantErr     error
	}{
		"適切なパラメータでメニューを作成できる": {
			id:          1,
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       150,
			stock:       100,
			wantErr:     nil,
		},
		"名前が空文字だとエラー": {
			id:          1,
			name:        "",
			description: "ほくほくでうまい",
			price:       150,
			stock:       100,
			wantErr:     domain.ErrInvalidName,
		},
		"名前が空白のみだとエラー": {
			id:          1,
			name:        " ",
			description: "ほくほくでうまい",
			price:       150,
			stock:       100,
			wantErr:     domain.ErrInvalidName,
		},
		"説明が空文字でもメニューを作成できる": {
			id:          1,
			name:        "かぼちゃ",
			description: "",
			price:       150,
			stock:       100,
			wantErr:     nil,
		},
		"priceが0でもメニューを作成できる": {
			id:          1,
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       0,
			stock:       100,
			wantErr:     nil,
		},
		"priceが負の数(-1)だとエラー": {
			id:          1,
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       -1,
			stock:       100,
			wantErr:     domain.ErrNegativePrice,
		},
		"stockが0でもメニューを作成できる": {
			id:          1,
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       150,
			stock:       0,
			wantErr:     nil,
		},
		"stockが負の数(-1)だとエラー": {
			id:          1,
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       150,
			stock:       -1,
			wantErr:     domain.ErrNegativeStock,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := domain.NewMenu(tt.id, tt.name, tt.description, tt.price, tt.stock)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err want: %v, got: %v", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				if got != nil {
					t.Errorf("Menu want: nil, got: %v", got)
				}
				return
			}
			if got.ID() != tt.id {
				t.Errorf("ID want: %d, got: %d", tt.id, got.ID())
			}
			if got.Name() != tt.name {
				t.Errorf("Name want: %v, got: %v", tt.name, got.Name())
			}
			if got.Description() != tt.description {
				t.Errorf("Description want: %v, got: %v", tt.description, got.Description())
			}
			if got.Price() != tt.price {
				t.Errorf("Price want: %d, got:%d", tt.price, got.Price())
			}
			if got.Stock() != tt.stock {
				t.Errorf("Stock want: %d, got: %d", tt.stock, got.Stock())
			}
		})
	}
}

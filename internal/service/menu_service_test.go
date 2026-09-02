package service_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/Suiren91/go-cafeteria/internal/service"
)

func TestCreateNewMenu(t *testing.T) {
	errStore := errors.New("store failure")
	tests := map[string]struct {
		name        string
		description string
		price       int
		stock       int

		createMenuFunc func(ctx context.Context, menu *domain.Menu) (int, error)

		wantID       int
		wantErr      error
		wantStoreHit bool
	}{
		"適切なパラメータでメニューを新規作成できる": {
			name:        "かぼちゃ",
			description: "ほくほくでうまい",
			price:       150,
			stock:       100,
			createMenuFunc: func(ctx context.Context, menu *domain.Menu) (int, error) {
				return 1, nil
			},
			wantID:       1,
			wantErr:      nil,
			wantStoreHit: true,
		}, "priceが負の数(-1)だとエラー": {
			name:           "かぼちゃ",
			description:    "ほくほくでうまい",
			price:          -1,
			stock:          100,
			createMenuFunc: nil,
			wantID:         0,
			wantErr:        domain.ErrNegativePrice,
			wantStoreHit:   false,
		}, "stockが負の数(-1)だとエラー": {
			name:           "かぼちゃ",
			description:    "ホクホクでうまい",
			price:          150,
			stock:          -1,
			createMenuFunc: nil,
			wantID:         0,
			wantErr:        domain.ErrNegativeStock,
			wantStoreHit:   false,
		}, "storeがエラーを返したらラップして返す": {
			name:        "かぼちゃ",
			description: "ほくほくで美味い",
			price:       150,
			stock:       100,
			createMenuFunc: func(ctx context.Context, menu *domain.Menu) (int, error) {
				return 0, errStore
			},
			wantID:       0,
			wantErr:      errStore,
			wantStoreHit: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mock := &MenuStoreMock{CreateMenuFunc: tt.createMenuFunc}
			s := service.NewMenuService(mock)

			gotID, err := s.CreateNewMenu(context.Background(),
				tt.name, tt.description, tt.price, tt.stock)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err want: %v, got: %v", tt.wantErr, err)
			}
			if gotID != tt.wantID {
				t.Errorf("id want: %v, got: %v", tt.wantID, gotID)
			}
			if hit := mock.CreateMenuCalls > 0; hit != tt.wantStoreHit {
				t.Errorf("store hit want: %v, got: %v", tt.wantStoreHit, hit)
			}
		})
	}
}

func TestListMenus(t *testing.T) {
	errStore := errors.New("store failure")
	menu := []*domain.Menu{
		domain.ReconstructMenu(1, "かぼちゃ", "ほくほくで美味い", 150, 100),
	}
	menus := []*domain.Menu{
		domain.ReconstructMenu(1, "かぼちゃ", "ほくほくで美味い", 150, 100),
		domain.ReconstructMenu(2, "にんじん", "あまい", 120, 50),
	}

	tests := map[string]struct {
		storeMenus []*domain.Menu
		storeErr   error

		wantMenus []*domain.Menu
		wantErr   error
	}{
		"保存されているメニューがスライスで返ってくる": {
			storeMenus: menus,
			storeErr:   nil,
			wantMenus:  menus,
			wantErr:    nil,
		}, "メニューが1件のときもスライスが返ってくる": {
			storeMenus: menu,
			storeErr:   nil,
			wantMenus:  menu,
			wantErr:    nil,
		}, "メニューが0件のときは空スライスが返ってくる": {
			storeMenus: []*domain.Menu{},
			storeErr:   nil,
			wantMenus:  []*domain.Menu{},
			wantErr:    nil,
		}, "storeがエラーを返したらラップして返す": {
			storeMenus: menus,
			storeErr:   errStore,
			wantMenus:  nil,
			wantErr:    errStore,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mock := &MenuStoreMock{
				ListMenusFunc: func(ctx context.Context) ([]*domain.Menu, error) {
					return tt.storeMenus, tt.storeErr
				},
			}
			s := service.NewMenuService(mock)

			gotMenus, err := s.ListMenus(context.Background())
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err want: %v, got: %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(gotMenus, tt.wantMenus) {
				t.Errorf("menus want: %v, got: %v", tt.wantMenus, gotMenus)
			}
			if mock.ListMenusCalls != 1 {
				t.Errorf("ListMenus calls want: 1, got: %v", mock.ListMenusCalls)
			}
		})
	}
}

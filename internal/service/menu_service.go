package service

import (
	"context"
	"fmt"

	"github.com/Suiren91/go-cafeteria/internal/domain"
)

type MenuStore interface {
	FindByID(ctx context.Context, id int) (*domain.Menu, error)
	CreateMenu(ctx context.Context, menu *domain.Menu) (int, error)
	ListMenus(ctx context.Context) ([]*domain.Menu, error)
}

type MenuService struct {
	store MenuStore
}

// NewMenuService は受け取ったMenuStoreを詰めた新しいMenuServiceを返す
func NewMenuService(s MenuStore) *MenuService {
	return &MenuService{store: s}
}

// CreateNewMenu は受け取った引数を使って新しいMenuを作る
func (s *MenuService) CreateNewMenu(ctx context.Context, name, description string, price, stock int) (int, error) {
	menu, err := domain.NewMenu(name, description, price, stock)
	if err != nil {
		return 0, err
	}
	id, err := s.store.CreateMenu(ctx, menu)
	if err != nil {
		return 0, fmt.Errorf("MenuService.CreateNewMenu: %w", err)
	}
	return id, nil
}

// ListMenus はDBに保存されているメニューを取得しスライスを返却する
func (s *MenuService) ListMenus(ctx context.Context) ([]*domain.Menu, error) {
	// NOTE: この関数のエラーがハンドラに渡ったとき，5xxエラーになるはず

	menus, err := s.store.ListMenus(ctx)
	if err != nil {
		return nil, fmt.Errorf("MenuService.ListMenus: %w", err)
	}
	return menus, nil
}

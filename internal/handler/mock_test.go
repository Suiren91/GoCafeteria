package handler_test

import (
	"context"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/Suiren91/go-cafeteria/internal/handler"
)

var _ handler.MenuService = (*MenuServiceMock)(nil)

type MenuServiceMock struct {
	CreateNewMenuFunc func(ctx context.Context, name, description string, price, stock int) (int, error)
	ListMenusFunc     func(ctx context.Context) ([]*domain.Menu, error)

	// CreateNewMenuCalls int
	// ListMenusCalls     int
}

func (m *MenuServiceMock) CreateNewMenu(ctx context.Context, name, description string, price, stock int) (int, error) {
	if m.CreateNewMenuFunc == nil {
		panic("MenuServiceMock: CreateNewMenuFunc is not set")
	}
	return m.CreateNewMenuFunc(ctx, name, description, price, stock)
}

func (m *MenuServiceMock) ListMenus(ctx context.Context) ([]*domain.Menu, error) {
	if m.ListMenusFunc == nil {
		panic("MenuServiceMock: ListMenusFunc is not set")
	}
	return m.ListMenusFunc(ctx)
}

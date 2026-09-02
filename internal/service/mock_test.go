package service_test

import (
	"context"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/Suiren91/go-cafeteria/internal/service"
)

var _ service.MenuStore = (*MenuStoreMock)(nil)

type MenuStoreMock struct {
	FindByIDFunc   func(ctx context.Context, id int) (*domain.Menu, error)
	CreateMenuFunc func(ctx context.Context, menu *domain.Menu) (int, error)
	ListMenusFunc  func(ctx context.Context) ([]*domain.Menu, error)

	CreateMenuCalls int
	ListMenusCalls  int
}

func (m *MenuStoreMock) FindByID(ctx context.Context, id int) (*domain.Menu, error) {
	if m.FindByIDFunc == nil {
		panic("MenuStoreMock: FindByIDFunc is not set")
	}
	return m.FindByIDFunc(ctx, id)
}

func (m *MenuStoreMock) CreateMenu(ctx context.Context, menu *domain.Menu) (int, error) {
	m.CreateMenuCalls++
	if m.CreateMenuFunc == nil {
		panic("MenuStoreMock: CreateMenuFunc is not set")
	}
	return m.CreateMenuFunc(ctx, menu)
}

func (m *MenuStoreMock) ListMenus(ctx context.Context) ([]*domain.Menu, error) {
	m.ListMenusCalls++
	if m.ListMenusFunc == nil {
		panic("MenuStoreMock: ListMenusFunc is not set")
	}
	return m.ListMenusFunc(ctx)
}

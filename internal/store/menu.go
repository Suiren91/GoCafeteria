package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Suiren91/go-cafeteria/internal/domain"
)

type MenuStore struct {
	q *Queries
}

// NewMenuStore は受け取ったQueriesを詰めた新しいMenuStoreを返す
func NewMenuStore(q *Queries) *MenuStore {
	return &MenuStore{q: q}
}

// FindByID は受け取ったidに一致するMenuをDBから検索し，見つかればそのMenuの構造体を返す
func (s *MenuStore) FindByID(ctx context.Context, id int) (*domain.Menu, error) {
	row, err := s.q.GetMenu(ctx, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrMenuNotFound
		}
		return nil, fmt.Errorf("store: get menu %d: %w", id, err)
	}
	menu := domain.ReconstructMenu(int(row.ID), row.Name, row.Description, int(row.PriceYen), int(row.Stock))
	return menu, nil
}

// CreateMenu は受け取ったMenuをDBに保存し，そのMenuのIDを返す
func (s *MenuStore) CreateMenu(ctx context.Context, menu *domain.Menu) (int, error) {
	int32Price := int32(menu.Price())
	if int(int32Price) != menu.Price() {
		return 0, fmt.Errorf("price out of int32 range: %d", menu.Price())
	}
	int32Stock := int32(menu.Stock())
	if int(int32Stock) != menu.Stock() {
		return 0, fmt.Errorf("stock out of int32 range: %d", menu.Stock())
	}
	menuParam := CreateMenuParams{
		Name:        menu.Name(),
		Description: menu.Description(),
		PriceYen:    int32Price,
		Stock:       int32Stock,
	}
	createdMenu, err := s.q.CreateMenu(ctx, menuParam)
	if err != nil {
		return 0, fmt.Errorf("store: create menu %q: %w", menu.Name(), err)
	}
	return int(createdMenu.ID), nil
}

// ListMenus はDBに保存されているMenuを呼び出し，スライスを返却する
func (s *MenuStore) ListMenus(ctx context.Context) ([]*domain.Menu, error) {
	smenus, err := s.q.ListMenus(ctx)
	if err != nil {
		return nil, fmt.Errorf("store: list menus: %w", err)
	}
	menus := make([]*domain.Menu, 0, len(smenus))
	for _, smenu := range smenus {
		m := domain.ReconstructMenu(
			int(smenu.ID),
			smenu.Name,
			smenu.Description,
			int(smenu.PriceYen),
			int(smenu.Stock),
		)
		menus = append(menus, m)
	}
	return menus, nil
}

// TODO: CountMenus
// TODO: UpdateMenu
// TODO: UpdateMenuPrice
// TODO: UpdateMenuStock

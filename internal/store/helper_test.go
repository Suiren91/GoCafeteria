package store_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/Suiren91/go-cafeteria/internal/store"
)

func newMenuStore(t *testing.T, ctx context.Context) *store.MenuStore {
	t.Helper()

	if testing.Short() {
		t.Skip("skipping test that requires a database (-short)")
	}
	// t.Context()はCleanupの直前にキャンセルされるため、
	// そのままトランザクションに紐づけるとCleanup内のRollbackが実行できない。
	tx, err := testPG.BeginTx(context.WithoutCancel(ctx), nil)
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}

	t.Cleanup(func() {
		if err := tx.Rollback(); err != nil &&
			!errors.Is(err, sql.ErrTxDone) {
			t.Errorf("rollback: %v", err)
		}
	})

	return store.NewMenuStore(store.New(tx))
}

type menuView struct {
	Name        string
	Description string
	Price       int
	Stock       int
}

func toView(m *domain.Menu) menuView {
	return menuView{
		Name:        m.Name(),
		Description: m.Description(),
		Price:       m.Price(),
		Stock:       m.Stock(),
	}
}

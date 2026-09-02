package store_test

import (
	"errors"
	"math"
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/google/go-cmp/cmp"
)

func TestMenuStore_CreateMenu_FindByID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		name        string
		description string
		price       int
		stock       int
	}{
		"通常のメニュー":  {"唐揚げ定食", "ジューシーな唐揚げ", 650, 10},
		"在庫0":      {"日替わり定食", "店員にお尋ねください", 550, 0},
		"説明が空文字":   {"わかめうどん", "", 300, 5},
		"マルチバイト文字": {"ラーメン🍜", "背脂チャッチャ系", 800, 3},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := t.Context()
			s := newMenuStore(t, ctx)

			want, err := domain.NewMenu(tt.name, tt.description, tt.price, tt.stock)
			if err != nil {
				t.Fatalf("NewMenu: %v", err)
			}
			id, err := s.CreateMenu(ctx, want)
			if err != nil {
				t.Fatalf("CreateMenu: %v", err)
			}
			if id < 0 {
				t.Fatalf("CreateMenu returned non-positive id: %d", id)
			}

			got, err := s.FindByID(ctx, id)
			if err != nil {
				t.Fatalf("FindByID(%d): %v", id, err)
			}
			if got.ID() != id {
				t.Errorf("ID: got: %d, want: %d", got.ID(), id)
			}
			if diff := cmp.Diff(toView(want), toView(got)); diff != "" {
				t.Errorf("menu mismatch (-want +got):\n%s ", diff)
			}
		})
	}
}

func TestMenuStore_FindByID_NotFound(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	s := newMenuStore(t, ctx)

	// 存在しないことが確実なID
	_, err := s.FindByID(ctx, math.MaxInt32)
	if !errors.Is(err, domain.ErrMenuNotFound) {
		t.Fatalf("got: %v, want: %v", err, domain.ErrMenuNotFound)
	}

	_, err = s.FindByID(ctx, -1)
	if !errors.Is(err, domain.ErrMenuNotFound) {
		t.Fatalf("got: %v, want: %v", err, domain.ErrMenuNotFound)
	}
}

func TestMenuStore_ListMenus(t *testing.T) {
	t.Parallel()

	t.Run("0件のときは空スライスを返す", func(t *testing.T) {
		t.Parallel()

		ctx := t.Context()
		s := newMenuStore(t, ctx)

		got, err := s.ListMenus(ctx)
		if err != nil {
			t.Fatalf("ListMenus: %v", err)
		}
		if got == nil {
			t.Error("got nil slice, want non-nil empty slice")
		}
		if len(got) != 0 {
			t.Errorf("got %d menus, want 0", len(got))
		}
	})

	t.Run("登録した全件を返す", func(t *testing.T) {
		t.Parallel()

		ctx := t.Context()
		s := newMenuStore(t, ctx)

		want := []menuView{
			{"唐揚げ定食", "ジューシー", 650, 10},
			{"にんじんステーキ", "つまみになる逸品", 350, 8},
			{"釜揚げうどん", "", 300, 5},
		}
		for _, w := range want {
			m, err := domain.NewMenu(w.Name, w.Description, w.Price, w.Stock)
			if err != nil {
				t.Fatalf("NewMenu(%q): %v", w.Name, err)
			}
			if _, err := s.CreateMenu(ctx, m); err != nil {
				t.Fatalf("CreateMenu(%q): %v", w.Name, err)
			}
		}

		menus, err := s.ListMenus(ctx)
		if err != nil {
			t.Fatalf("ListMenus: %v", err)
		}

		got := make([]menuView, 0, len(menus))
		for _, m := range menus {
			got = append(got, toView(m))
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("menus mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMenuStore_CreateMenu_OutOfInt32Range(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	s := newMenuStore(t, ctx)

	menu := domain.ReconstructMenu(0, "壊れたメニュー", "特に値段が壊れている", math.MaxInt32+1, 1)

	if _, err := s.CreateMenu(ctx, menu); err == nil {
		t.Fatal("want error for out-of-int32 price, got nil")
	}
}

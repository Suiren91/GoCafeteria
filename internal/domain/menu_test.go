package domain_test

import (
	"testing"

	"go-cafeteria/internal/domain"
)

func TestNewMenu(t *testing.T) {
	t.Run("正常に新規作成できる", func(t testing.T) {
		m, err = domain.NewMenu("卵焼き", "ほくほくでとろとろ，美味しい卵焼き", 100, 5)
		if err != nil {
			t.Errorf("後で直す")
		}
	})
}

package domain_test

import (
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/domain"
)

const (
	ID          int    = 1
	NAME        string = "かぼちゃ"
	DESCRIPTION string = "ほくほくでうまい"
	PRICE       int    = 150
	STOCK       int    = 100
)

func TestNewMenu(t *testing.T) {
	got, err := domain.NewMenu(ID, NAME, DESCRIPTION, PRICE, STOCK)
	if err != nil {
		t.Fatalf("NewMenu() returned unexpected error: %v", err)
	}
	if got.ID() != ID {
		t.Errorf("ID want: %d, got: %d", ID, got.ID())
	}
	if got.Name() != NAME {
		t.Errorf("Name want: %v, got: %v", NAME, got.Name())
	}
	if got.Description() != DESCRIPTION {
		t.Errorf("Description want: %v, got: %v", DESCRIPTION, got.Description())
	}
	if got.Price() != PRICE {
		t.Errorf("Price want: %d, got:%d", PRICE, got.Price())
	}
	if got.Stock() != STOCK {
		t.Errorf("Stock want: %d, got: %d", STOCK, got.Stock())
	}
}

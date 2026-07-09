package service_test

import (
	"testing"

	"github.com/Suiren91/go-cafeteria/internal/service"
)

const (
	ID          int    = 1
	NAME        string = "かぼちゃ"
	DESCRIPTION string = "ほくほくでうまい"
	PRICE       int    = 150
	STOCK       int    = 100
)

func TestCreateNewMenu(t *testing.T) {
	err := service.CreateNewMenu(ID, NAME, DESCRIPTION, PRICE, STOCK)
	if err != nil {
		t.Errorf("TestCreateNewMenu failed: %v", err)
	}
}

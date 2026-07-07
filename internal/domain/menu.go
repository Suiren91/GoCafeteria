// Package domainは業務で必要なエンティティや値オブジェクトを提供します
package domain

import (
	"errors"
)

type Menu struct {
	id          int
	name        string
	description string
	price       int
	stock       int
}

// NewMenu は新しいMenuを生成して返します．priceかstockが0未満であればerrorを返します
func NewMenu(id int, name, description string, price, stock int) (*Menu, error) {
	if price < 0 {
		return &Menu{}, errors.New("price must be non-negative")
	}
	if stock < 0 {
		return &Menu{}, errors.New("stock must be non-negative")
	}
	return &Menu{
		id:          id,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func (m *Menu) ID() int {
	return m.id
}

func (m *Menu) Name() string {
	return m.name
}

func (m *Menu) Description() string {
	return m.description
}

func (m *Menu) Price() int {
	return m.price
}

func (m *Menu) Stock() int {
	return m.stock
}

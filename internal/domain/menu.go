// Package domainは業務で必要なエンティティや値オブジェクトを提供します
package domain

import (
	"strings"
)

type Menu struct {
	id          int
	name        string
	description string
	price       int
	stock       int
}

// NewMenu は新しいMenuを生成して返します．
// 以下の場合にエラーを返します:
//   - nameが空文字か空白のみのとき: [ErrInvalidName]
//   - priceが負のとき: [ErrNegativePrice]
//   - stockが負のとき: [ErrNegativeStock]
func NewMenu(id int, name, description string, price, stock int) (*Menu, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidName
	}
	if price < 0 {
		return nil, ErrNegativePrice
	}
	if stock < 0 {
		return nil, ErrNegativeStock
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

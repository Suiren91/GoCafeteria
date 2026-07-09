package service

import "github.com/Suiren91/go-cafeteria/internal/domain"

// TODO: 構造体とコンストラクタ
func CreateNewMenu(id int, name, description string, price, stock int) error {
	_, err := domain.NewMenu(id, name, description, price, stock)
	// TODO: DBへの保存処理
	if err != nil {
		return err
	}
	return nil
}

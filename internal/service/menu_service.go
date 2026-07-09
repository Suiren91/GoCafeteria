package service

import "github.com/Suiren91/go-cafeteria/internal/domain"

// TODO: id管理の方法を模索
var id int = 1

// TODO: 構造体とコンストラクタ
func CreateNewMenu(name, description string, price, stock int) error {
	_, err := domain.NewMenu(id, name, description, price, stock)
	id++
	// TODO: DBへの保存処理
	if err != nil {
		// TODO: 適切なエラーに翻訳して返却
		return err
	}
	return nil
}

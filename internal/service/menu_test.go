package service_test

import (
	"context"
	"testing"
)

func TestNewMenu(t *testing.T){
	t.Run("新しいメニューを作成できる", f func(t *testing.T){
		menu:=&mockMenu{
//menuのフィールドを描く
		}
		svc := NewMenuService()

	})
}

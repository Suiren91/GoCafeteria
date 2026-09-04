package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"uuid"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	guuid "github.com/google/uuid"
)

type UserPasswordStore struct {
	q *Queries
}

// NewUserPasswordStore はユーザーパスワードを保存するための
// NewUserPasswordStore構造体を生成する
func NewUserPasswordStore(q *Queries) *UserPasswordStore {
	return &UserPasswordStore{q: q}
}

// CreatePassword はユーザーのパスワードをデータベースに保存する
// 処理の途中で発生したエラーはラップして返す
func (s *UserPasswordStore) CreatePassword(ctx context.Context, id uuid.UUID, hashPw string) error {
	param := CreatePasswordHashParams{
		UserID:       guuid.UUID(id),
		PasswordHash: hashPw,
	}
	_, err := s.q.CreatePasswordHash(ctx, param)
	if err != nil {
		return fmt.Errorf("store: create password: %w", err)
	}
	return nil
}

// GetPassword はユーザーのパスワードをデータベースから取得する
// 処理の途中で発生したエラーはラップして返す
func (s *UserPasswordStore) GetPassword(ctx context.Context, id uuid.UUID) (string, error) {
	pw, err := s.q.GetPasswordHash(ctx, guuid.UUID(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// WARNING: このコードを通るのはユーザーが見つからなかったにも関わらず
			// WARNING: GetPasswordを呼び出した時か，もしくはユーザーは保存できたが
			// WARNING: パスワードが保存できなかった場合のみ
			// WARNING: つまり，整合性が取れていない異常状態のため，通常とは異なる対応が必要
			return "", domain.ErrUserNotFound
		}
		return "", fmt.Errorf("store: get password: %w", err)
	}
	return pw, nil
}

// UpdatePassword は指定したユーザーのパスワードを変更する
// 処理の途中で発生したエラーはラップして返す
func (s *UserPasswordStore) UpdatePassword(ctx context.Context, id uuid.UUID, hashPw string) error {
	param := UpdatePasswordHashParams{
		UserID:       guuid.UUID(id),
		PasswordHash: hashPw,
	}
	_, err := s.q.UpdatePasswordHash(ctx, param)
	if err != nil {
		fmt.Errorf("store: update password: %w", err)
	}
	return nil
}

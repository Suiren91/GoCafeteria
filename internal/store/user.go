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

type UserStore struct {
	q *Queries
}

// NewUserStore はUserをデータベースに保存するための
// UserStore構造体を生成します
func NewUserStore(q *Queries) *UserStore {
	return &UserStore{q: q}
}

// FindByEmail は与えられたemailに合致するユーザーを返す
// ユーザーが見つからなければ[ErrUserNotFound]を返す
// そのほかのデータベースで起きたエラーはラップして返す
func (s *UserStore) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	rowUser, err := s.q.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("store: get user: %w", err)
	}
	return domain.ReconstructUser(uuid.UUID(rowUser.ID), rowUser.Email, rowUser.Name)
}

// ListUsers はデータベースに保存されているユーザーのスライスを返す
// 処理の途中で発生したエラーはラップして返す
func (s *UserStore) ListUsers(ctx context.Context) ([]*domain.User, error) {
	rows, err := s.q.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("store: list users: %w", err)
	}
	users := make([]*domain.User, 0, len(rows))
	for _, row := range rows {
		user, err := domain.ReconstructUser(uuid.UUID(row.ID), row.Email, row.Name)
		if err != nil {
			return nil, fmt.Errorf("store: list users: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

// CreateUser はデータベースにユーザーを登録する
// 処理の途中で発生したエラーはラップして返す
func (s *UserStore) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	param := CreateUserParams{
		ID:    guuid.UUID(user.ID()),
		Email: user.Email().String(),
		Name:  user.Name(),
	}

	row, err := s.q.CreateUser(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("store: create user: %w", err)
	}

	reUser, err := domain.ReconstructUser(uuid.UUID(row.ID), row.Email, row.Name)
	if err != nil {
		return nil, fmt.Errorf("store: create user: %w", err)
	}
	return reUser, nil
}

// UpdateUser はデータベースに保存されているデータを更新する
// 更新しない情報はそのまま引数に渡してください
// 処理の途中で発生したエラーはラップして返す
func (s *UserStore) UpdateUser(ctx context.Context, id uuid.UUID, email domain.Email, name string) (*domain.User, error) {
	param := UpdateUserParams{
		ID:    guuid.UUID(id),
		Email: email.String(),
		Name:  name,
	}

	row, err := s.q.UpdateUser(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("store: update user: %w", err)
	}
	reUser, err := domain.ReconstructUser(uuid.UUID(row.ID), row.Email, row.Name)
	if err != nil {
		return nil, fmt.Errorf("store: update user: %w", err)
	}
	return reUser, nil
}

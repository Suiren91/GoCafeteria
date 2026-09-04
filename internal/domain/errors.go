package domain

import "errors"

var (
	ErrNegativePrice = errors.New("price must be non-negative")
	ErrNegativeStock = errors.New("stock must be non-negative")
	ErrInvalidID     = errors.New("id must be positive")
	ErrInvalidName   = errors.New("name must not be empty")
	ErrMenuNotFound  = errors.New("menu not found")
	ErrInvalidEmail  = errors.New("invalid email address")
	ErrUserNotFound  = errors.New("user not found")
)

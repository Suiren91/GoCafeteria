package domain

import (
	"net/mail"
	"strings"
	"uuid"
)

type User struct {
	id    uuid.UUID
	email Email
	//ユーザーの表示名
	//空白のみも許容する(空文字は許容しない)
	name string
}

// Email型の定義は他のコードからも参照するようになったら
// email.goとして別コードに切り出す
type Email struct {
	value string
}

// newEmail は与えられたメールアドレスを検証し，小文字に変換した上で
// Email構造体として返却する
// メールアドレスとして不正な形であればエラーを返す
// 試しにunexportedにしてる(外で使いたくなったらexportする)
func newEmail(s string) (Email, error) {
	addr, err := mail.ParseAddress(s)
	if err != nil {
		return Email{}, ErrInvalidEmail
	}
	address := strings.ToLower(addr.Address)
	return Email{value: address}, nil
}

func (e Email) String() string {
	return e.value
}

// NewUser は与えられたemail, nameを持つ新しいUser構造体を作成する
// emailが不正なときは[ErrInvalidEmail]を返す
// Nameが空文字("")のときは[ErrInvalidName]を返す
func NewUser(email, name string) (*User, error) {
	validEmail, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	if name == "" {
		return nil, ErrInvalidName
	}
	return &User{
		email: validEmail,
		name:  name,
	}, nil
}

// ReconstructUserは永続化されたデータからUserを復元する
// emailが不正なときは[ErrInvalidEmail]を返す
func ReconstructUser(id uuid.UUID, email, name string) (*User, error) {
	// 本来ReconsstructUserからエラーを返したくないが，Emailが独自型のため，
	// 不正なEmailはエラーを返す
	validEmail, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	return &User{
		id:    id,
		email: validEmail,
		name:  name,
	}, nil
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) Name() string {
	return u.name
}

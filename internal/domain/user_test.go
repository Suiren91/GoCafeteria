package domain_test

import (
	"errors"
	"testing"
	"uuid"

	"github.com/Suiren91/go-cafeteria/internal/domain"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		email     string
		name      string
		wantEmail string
		wantErr   error
	}{
		"正しいメールアドレスと表示名でユーザーを作成できる": {
			email:     "hoge@example.com",
			name:      "hoge",
			wantEmail: "hoge@example.com",
			wantErr:   nil,
		},
		"名前が空白のみでもユーザーを作成できる": {
			email:     "fuga@example.com",
			name:      " ",
			wantEmail: "fuga@example.com",
			wantErr:   nil,
		},
		"表示名形式でメールアドレスを渡されてもアドレスのみを取り出してユーザーを作成できる": {
			email:     "Taro Tanaka<taro@example.com>",
			name:      "Taro",
			wantEmail: "taro@example.com",
			wantErr:   nil,
		},
		"大文字小文字の混じったメールアドレスはすべて小文字に正規化される": {
			email:     "TaroTanaka@example.com",
			name:      "Taro",
			wantEmail: "tarotanaka@example.com",
			wantErr:   nil,
		},
		"名前が日本語(2バイト文字)でもユーザーを作成できる": {
			email:     "taro@example.com",
			name:      "田中太郎",
			wantEmail: "taro@example.com",
			wantErr:   nil,
		},
		"名前が空文字だとエラー": {
			email:     "hoge@example.com",
			name:      "",
			wantEmail: "",
			wantErr:   domain.ErrInvalidName,
		},
		"メールアドレスが空文字だとエラー": {
			email:     "",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"メールアドレスが空白のみだとエラー": {
			email:     " ",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"空白の入ったメールアドレスだとエラー": {
			email:     "hoge @example.com",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"ローカル部が存在しないメールアドレスだとエラー": {
			email:     "@example.com",
			name:      "home",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"ドメイン部が存在しないメールアドレスだとエラー": {
			email:     "hoge@",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"@が存在しないメールアドレスだとエラー": {
			email:     "hogehogeexample.com",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewUser(tt.email, tt.name)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewUser: err: want: %v, got: %v", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				if got != nil {
					t.Errorf("NewUser: user: want: nil, got: %v", got)
				}
				return
			}

			if got.Email().String() != tt.wantEmail {
				t.Errorf("NewUser: Email: want: %v, got: %v", tt.wantEmail, got.Email().String())
			}
			if got.Name() != tt.name {
				t.Errorf("NewUser: Name: want: %v, got: %v", tt.name, got.Name())
			}
		})
	}
}

func TestReconstructUser(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id    uuid.UUID
		email string
		name  string

		wantEmail string
		wantErr   error
	}{
		"正常にユーザーを復元できる": {
			id:        uuid.NewV7(),
			email:     "hoge@example.com",
			name:      "hoge",
			wantEmail: "hoge@example.com",
			wantErr:   nil,
		},
		"名前が空白のみでもユーザーを復元できる": {
			id:        uuid.NewV7(),
			email:     "fuga@example.com",
			name:      " ",
			wantEmail: "fuga@example.com",
			wantErr:   nil,
		},
		"表示名形式でメールアドレスを渡されてもアドレスのみを取り出してユーザーを復元できる": {
			id:        uuid.NewV7(),
			email:     "Taro Tanaka<taro@example.com>",
			name:      "Taro",
			wantEmail: "taro@example.com",
			wantErr:   nil,
		},
		"大文字小文字の混じったメールアドレスはすべて小文字に正規化される": {
			id:        uuid.NewV7(),
			email:     "TaroTanaka@example.com",
			name:      "Taro",
			wantEmail: "tarotanaka@example.com",
			wantErr:   nil,
		},
		"名前が日本語(2バイト文字)でもユーザーを復元できる": {
			id:        uuid.NewV7(),
			email:     "taro@example.com",
			name:      "田中太郎",
			wantEmail: "taro@example.com",
			wantErr:   nil,
		},
		"名前が空文字でもユーザーを復元できる": {
			id:        uuid.NewV7(),
			email:     "hoge@example.com",
			name:      "",
			wantEmail: "hoge@example.com",
			wantErr:   nil,
		},
		"メールアドレスが空文字だとエラー": {
			id:        uuid.NewV7(),
			email:     "",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"メールアドレスが空白のみだとエラー": {
			id:        uuid.NewV7(),
			email:     " ",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"空白の入ったメールアドレスだとエラー": {
			id:        uuid.NewV7(),
			email:     "hoge @example.com",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"ローカル部が存在しないメールアドレスだとエラー": {
			id:        uuid.NewV7(),
			email:     "@example.com",
			name:      "home",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"ドメイン部が存在しないメールアドレスだとエラー": {
			id:        uuid.NewV7(),
			email:     "hoge@",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
		"@が存在しないメールアドレスだとエラー": {
			id:        uuid.NewV7(),
			email:     "hogehogeexample.com",
			name:      "hoge",
			wantEmail: "",
			wantErr:   domain.ErrInvalidEmail,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.ReconstructUser(tt.id, tt.email, tt.name)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ReconstructUser: err: want: %v, got: %v", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				if got != nil {
					t.Errorf("ReconstructUser: user: want: nil, got: %v", got)
				}
				return
			}

			if got.ID() != tt.id {
				t.Errorf("ReconstructUser: ID: want: %v, got: %v", tt.id, got.ID())
			}
			if got.Email().String() != tt.wantEmail {
				t.Errorf("ReconstructUser: Email: want: %v, got: %v", tt.wantEmail, got.Email().String())
			}
			if got.Name() != tt.name {
				t.Errorf("ReconstructUser: Name: want: %v, got: %v", tt.name, got.Name())
			}
		})
	}
}

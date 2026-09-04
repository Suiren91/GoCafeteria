CREATE TABLE users (
  id uuid PRIMARY KEY ,
  email text NOT NULL UNIQUE,
  name text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON COLUMN users.name IS '表示名';

CREATE TABLE user_passwords(
user_id uuid PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
password_hash text NOT NULL,
updated_at timestamptz NOT NULL DEFAULT now()
);

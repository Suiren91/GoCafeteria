-- name: GetPasswordHash :one
SELECT password_hash FROM user_passwords WHERE user_id = $1;

-- name: CreatePasswordHash :one
INSERT INTO user_passwords (user_id, password_hash) VALUES($1,$2) RETURNING *;
-- name: UpdatePasswordHash :one
UPDATE user_passwords SET password_hash=$2, updated_at=now() WHERE user_id = $1 RETURNING *;


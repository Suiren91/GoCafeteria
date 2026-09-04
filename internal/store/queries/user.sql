-- name: GetUser :one
SELECT * FROM users WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: CreateUser :one
INSERT INTO users(id, email, name) VALUES ($1,$2,$3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET email=$2, name=$3, updated_at=now() WHERE id=$1 RETURNING *;

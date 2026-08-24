-- name: GetMenu :one
SELECT * FROM menus WHERE id = $1;

-- name: ListMenus :many
SELECT * FROM menus ORDER BY id;

-- name: CountMenus :one
SELECT count(*) FROM menus;

-- name: CreateMenu :one
INSERT INTO menus (name, description, price_yen, stock) VALUES($1,$2,$3,$4) RETURNING *;

--name UpdateMenu :one
UPDATE menus SET name=$2, description=$3, price_yen=$4, stock=$5 WHERE id=$1 RETURNING *;

--name UpdateMenuPrice :one
UPDATE menus SET price_yen=$2 WHERE id=$1 RETURNING *;

--name UpdateMenuStock :one
UPDATE menus SET stock=$2 WHERE id=$1 RETURNING *;


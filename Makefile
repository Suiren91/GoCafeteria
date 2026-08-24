include .env
export

DATABASE_URL = postgres://user:${PG_PASSWORD}@localhost:5432/cafeteria?sslmode=disable


up: 
	docker compose up

pg-up:
	docker compose up -d --wait db

migrate-up:
	go run -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate \
		-path internal/store/migrations -database "$(DATABASE_URL)" up

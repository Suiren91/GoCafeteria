include .env
export

PG_HOST ?= localhost
PG_PORT ?= 5432
PG_TEST_PORT ?= 5433
PG_USER ?= user
PG_DB ?= cafeteria
PG_TEST_DB ?= cafeteria_test

PG_URL = postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable
PG_TEST_URL = postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_TEST_PORT)/$(PG_TEST_DB)?sslmode=disable

TEST_ENV = PG_PORT=$(PG_TEST_PORT) PG_DB=$(PG_TEST_DB)

up: 
	docker compose up

pg-up:
	docker compose up -d --wait rdb

pg-test-up:
	docker compose --profile test up -d --wait rdb-test

#test-db: pg-up
#	@docker compose exec -T rdb psql -U $(PG_USER) -d $(PG_DB) -tAc \
#		"SELECT 1 FROM pg_database WHERE datname='$(PG_TEST_DB)" | grep -q 1 \
#		|| docker compose exec -T rdb createdb -U $(PG_USER) $(PG_TEST_DB)

migrate-up:
	go run -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate \
		-path internal/store/migrations -database "$(PG_URL)" up

coverage: pg-test-up
	$(TEST_ENV) go test -v -coverprofile=coverage.out ./...

coverout:
	go tool cover -html=coverage.out -o coverage.html

test: pg-test-up
	$(TEST_ENV) go test ./...

stest:
	go test -short ./...

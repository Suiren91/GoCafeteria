package store_test

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	pgxmigrate "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var testPG *sql.DB

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Store layer test setup failed: ", err)
		os.Exit(1)
	}
	os.Exit(code)
}

func run(m *testing.M) (int, error) {
	flag.Parse()

	if testing.Short() {
		return m.Run(), nil
	}

	if err := godotenv.Load("../../.env"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return 0, fmt.Errorf("godotenv: %w", err)
	}

	// TODO: configパッケージで設定するようにするかも?
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_DB"))

	pg, err := sql.Open("pgx/v5", url)
	if err != nil {
		return 0, fmt.Errorf("sql.Open: %w", err)
	}
	defer func() { _ = pg.Close() }()

	if err := pg.Ping(); err != nil {
		return 0, fmt.Errorf("ping: %w", err)
	}
	if err := applyMigrations(pg); err != nil {
		return 0, fmt.Errorf("apply migrations: %w", err)
	}
	testPG = pg
	return m.Run(), nil
}

func applyMigrations(pg *sql.DB) error {
	driver, err := pgxmigrate.WithInstance(pg, &pgxmigrate.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}

	if err := mig.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

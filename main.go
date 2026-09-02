package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Suiren91/go-cafeteria/internal/handler"
	"github.com/Suiren91/go-cafeteria/internal/service"
	"github.com/Suiren91/go-cafeteria/internal/store"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		slog.Error("startup failed", "err", err)
		os.Exit(1)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("godotenv.Load: %w", err)
	}

	rootCtx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 依存関係の構築
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_DB"))
	pg, err := sql.Open("pgx", url)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	defer pg.Close()

	pingCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()
	if err := pg.PingContext(pingCtx); err != nil {
		return fmt.Errorf("conn.PingContext: %w", err)
	}

	q := store.New(pg)
	menuHdl := handler.NewMenuHandler(
		service.NewMenuService(
			store.NewMenuStore(q),
		),
	)
	r := handler.NewRouter(menuHdl)

	// TODO: 環境変数からポートを取るようにする？
	srv := &http.Server{Addr: ":8080", Handler: r}
	// TODO: graceful shutdown
	return srv.ListenAndServe()
}

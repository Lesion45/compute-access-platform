package postgres

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Postgres struct {
	DB *pgxpool.Pool
}

func NewPostgres(ctx context.Context, dsn string) *Postgres {
	var (
		pgInstance *Postgres
		pgOnce     sync.Once
	)

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, dsn)
		if err != nil {
			log.Fatal("Unable to connect to database", zap.Error(err))
		}

		pgInstance = &Postgres{
			DB: db,
		}
	})
	return pgInstance
}

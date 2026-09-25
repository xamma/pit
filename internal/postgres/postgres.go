package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Init(ctx context.Context, connString, sqlFile string) error {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	sql, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, string(sql))
	return err
}

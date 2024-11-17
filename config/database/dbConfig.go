package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func InitDatabase(db_url string) error {
	if db_url == "" {
		return fmt.Errorf("DB url is not configured")
	}

	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	return nil
}

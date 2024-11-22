package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func InitDatabase(db_url string) (*pgx.Conn, error) {
	if db_url == "" {
		return nil, fmt.Errorf("DB url is not configured")
	}

	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

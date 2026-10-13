package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type Query struct {
	DBTX DBTX
}

func (q *Query) GetItemByID(ctx context.Context, itemID int) pgx.Row {
	return q.DBTX.QueryRow(ctx, `
		SELECT id, name, description, price, stock, created_at
		FROM items
		WHERE id = $1
	`, itemID)
}

func (q *Query) GetItems(ctx context.Context) (pgx.Rows, error) {
	return q.DBTX.Query(ctx, `
		SELECT id, name, description, price, stock, created_at
		FROM items
	`)
}

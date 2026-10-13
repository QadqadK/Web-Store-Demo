package services

import (
	"context"
	"errors"
	"fmt"
	"webstore/models"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	conn *pgx.Conn
}

func NewStore(conn *pgx.Conn) *Store {
	return &Store{
		conn: conn,
	}
}

func (s *Store) Queries() *Query {
	return &Query{
		DBTX: s.conn,
	}
}

func (s *Store) WithTx(tx pgx.Tx) *Query {
	return &Query{
		DBTX: tx,
	}
}

func (s *Store) GetItem(ctx context.Context, itemID int) (*models.Item, error) {
	var item models.Item
	err := s.Queries().GetItemByID(ctx, itemID).Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Stock, &item.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (s *Store) GetItems(ctx context.Context) ([]*models.Item, error) {
	rows, err := s.Queries().GetItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to select all items", err)
	}
	defer rows.Close()

	items := make([]*models.Item, 0)
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Stock, &item.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to scan item row", err)
		}
		items = append(items, &item)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return items, nil
}

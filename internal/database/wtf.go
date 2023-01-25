package database

import (
	"context"
	"test/test/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type WtfService struct {
	*sqlx.DB
}

func (w *WtfService) Wtfs(filter domain.WtfFilter) ([]*domain.WTF, int, error) {

	// TODO: fake query
	query := `
	SELECT *
	FROM wtfs`

	// TODO: fake args
	// args := []any{filter.IDs, filter.Wtfs, filter.Limit, filter.Offset}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := w.DB.QueryContext(ctx, query) // args...
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	wtfs := []*domain.WTF{}

	for rows.Next() {
		var wtf domain.WTF

		err := rows.Scan(
			&wtf.ID,
			&wtf.Wtf,
			&wtf.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		wtfs = append(wtfs, &wtf)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return wtfs, len(wtfs), nil
}

func (w *WtfService) WtfByID(id int) (*domain.WTF, error) {
	query := `SELECT * FROM wtfs WHERE id = $1;`

	wtf := domain.WTF{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := w.DB.QueryRowContext(ctx, query, id).Scan(&wtf.ID, &wtf.Wtf, &wtf.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &wtf, nil
}

func (w *WtfService) CreateWtf(wtf *domain.WTF) error {
	return nil
}
func (w *WtfService) UpdateWtf(id int, wtf *domain.WtfUpdate) (*domain.WTF, error) {
	return nil, nil
}
func (w *WtfService) DeleteWtf(id int) error {
	return nil
}

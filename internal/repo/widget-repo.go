package repo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Widget struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type WidgetRepo struct {
	pool *pgxpool.Pool
}

func NewWidgetRepo(pool *pgxpool.Pool) *WidgetRepo {
	return &WidgetRepo{pool: pool}
}

func (r *WidgetRepo) List(ctx context.Context) ([]Widget, error) {
	rows, err := r.pool.Query(ctx, `
		select id, name, created_at
		from widgets
		order by id asc
		limit 100
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Widget
	for rows.Next() {
		var w Widget
		if err := rows.Scan(&w.ID, &w.Name, &w.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, w)
	}
	return out, rows.Err()
}

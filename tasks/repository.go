package tasks

import (
	"context"
	"database/sql"
)

type Repository interface {
	FindAll(ctx context.Context) (Task, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
    return &repository{
        db: db,
    }
}
// FindAll implements Repository.
func (r *repository) FindAll(ctx context.Context) (Task, error) {
    result := Task{}
    err := r.db.QueryRowContext(ctx, `
        SELECT
            id,
            title,
            description,
            completed
        FROM
            tasks
    `).Scan(
            &result.ID,
            &result.Title,
            &result.Description,
            &result.Completed,
        )
        if err != nil {
            if err == sql.ErrNoRows {
                return Task{}, sql.ErrNoRows
            }
            return Task{}, err
        }
        return result, nil
    }


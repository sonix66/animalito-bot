package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sonix66/animalito-bot/internal/entity"
)

func (r *Repository) CreatePhoto(
	ctx context.Context,
	animalID string,
	callback entity.OnSaveCallback,
) error {
	err := r.runInTransaction(ctx, func(context.Context, *sql.Tx) error {
		id := uuid.New().String()

		err := callback(id)
		if err != nil {
			return fmt.Errorf("callback: %w", err)
		}

		queryPhoto := `
			INSERT INTO photos (id, animal_id)
			VALUES ($1, $2)
		`
		_, err = r.db.ExecContext(ctx, queryPhoto, id, animalID)
		if err != nil {
			return fmt.Errorf("r.db.ExecContext: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("r.runInTransaction: %w", err)
	}

	return nil

}

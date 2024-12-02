package sqlite

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (r *Repository) DeleteAnimalByID(ctx context.Context, id string) error {
	if id == "" {
		return entity.ErrEmptyID
	}

	query := `
		DELETE FROM animals
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("r.db.ExecContext: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected: %w", err)
	}

	return nil
}

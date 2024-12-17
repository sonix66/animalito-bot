package sqlite

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (r *Repository) UpdateAnimal(
	ctx context.Context,
	animal *entity.Animal,
) error {
	if animal.ID == "" {
		return entity.ErrEmptyID
	}

	query := `
		UPDATE animals
		SET name = $1,
			type = $2,
			description = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query,
		animal.Name,
		animal.Type,
		animal.Description,
		animal.ID,
	)
	if err != nil {
		return fmt.Errorf("r.db.ExecContext: %w", err)
	}

	return nil
}

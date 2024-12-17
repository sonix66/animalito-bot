package sqlite

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sonix66/animalito-bot/internal/entity"
)

func (r *Repository) CreateAnimal(ctx context.Context, animal *entity.Animal) (string, error) {
	animal.ID = uuid.NewString()

	// Запрос для вставки животного
	queryAnimal := `
		INSERT INTO animals (id, name, type, description)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, queryAnimal,
		animal.ID,
		animal.Name,
		animal.Type,
		animal.Description,
	)
	if err != nil {
		return "", fmt.Errorf("r.db.QueryRowContext: %w", err)
	}

	return animal.ID, nil
}

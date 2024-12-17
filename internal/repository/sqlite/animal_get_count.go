package sqlite

import (
	"context"
	"fmt"
)

func (r *Repository) GetAnimalsCount(ctx context.Context) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM animals
	`

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("r.db.QueryRowContext: %w", err)
	}

	return count, nil
}

package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (r *Repository) GetAnimalByID(ctx context.Context, id string) (*entity.Animal, error) {
	if id == "" {
		return nil, entity.ErrEmptyID
	}

	query := `
		SELECT a.id, a.name, a.type, a.description, a.created_at,
		       COALESCE(GROUP_CONCAT(p.id), '') AS photo_ids
		FROM animals a
		LEFT JOIN photos p ON a.id = p.animal_id
		WHERE a.id = $1
		GROUP BY a.id
	`

	var animal entity.Animal
	var photoIDs sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&animal.ID,
		&animal.Name,
		&animal.Type,
		&animal.Description,
		&animal.CreatedAt,
		&photoIDs,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, fmt.Errorf("r.db.QueryRowContext: %w", err)
	}

	if photoIDs.Valid {
		animal.PhotoURLs = splitPhotoIDs(photoIDs.String)
	} else {
		animal.PhotoURLs = []string{}
	}

	return &animal, nil
}

func splitPhotoIDs(photoIDs string) []string {
	if photoIDs == "" {
		return []string{}
	}
	return strings.Split(photoIDs, ",")
}

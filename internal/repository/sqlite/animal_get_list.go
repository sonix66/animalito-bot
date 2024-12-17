package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (r *Repository) GetAnimalList(
	ctx context.Context,
	count, offset int,
) ([]*entity.Animal, error) {
	query := `
		SELECT a.id, a.name, a.type, a.description, a.created_at,
		       COALESCE(GROUP_CONCAT(p.id), '') AS photo_ids
		FROM animals a
		LEFT JOIN photos p ON a.id = p.animal_id
		GROUP BY a.id
		ORDER BY a.created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, count, offset)
	if err != nil {
		return nil, fmt.Errorf("r.db.QueryContext: %w", err)
	}
	defer rows.Close()

	var animals []*entity.Animal

	for rows.Next() {
		var animal entity.Animal
		var photoIDs sql.NullString

		errScan := rows.Scan(
			&animal.ID,
			&animal.Name,
			&animal.Type,
			&animal.Description,
			&animal.CreatedAt,
			&photoIDs,
		)
		if errScan != nil {
			return nil, fmt.Errorf("rows.Scan: %w", errScan)
		}

		if photoIDs.Valid {
			animal.PhotoURLs = splitPhotoIDs(photoIDs.String)
		} else {
			animal.PhotoURLs = []string{}
		}

		animals = append(animals, &animal)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return animals, nil
}

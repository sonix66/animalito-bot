package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) runInTransaction(ctx context.Context, f func(context.Context, *sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("r.db.Begin: %w", err)
	}

	err = f(ctx, tx)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("tx.Rollback: %w", rollbackErr)
		}

		return fmt.Errorf("f: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}

	return nil
}

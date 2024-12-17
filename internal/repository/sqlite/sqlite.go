package sqlite

import (
	"database/sql"
)

type Config struct {
	FilePath string `mapstructure:"filepath"`
}

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

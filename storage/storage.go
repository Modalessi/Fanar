package storage

import (
	"database/sql"

	"github.com/Modalessi/iau_resources/database"
)

type Storage struct {
	db      *sql.DB
	queries *database.Queries
}

func NewStorage(db *sql.DB, queries *database.Queries) *Storage {
	return &Storage{
		db:      db,
		queries: queries,
	}
}

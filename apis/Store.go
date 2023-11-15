package apis

import (
	"database/sql"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
)

type Store struct {
	db *sql.DB
	*db.Queries
}

func NewStore(dbs *sql.DB) *Store {
	return &Store{
		db:      dbs,
		Queries: db.New(dbs),
	}
}

func (s *Store) DBTransaction() (*sql.Tx, error) {
	return s.db.Begin()
}

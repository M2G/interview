package db

import (
	"context"
	"database/sql"
)

type Queries struct {
	db *sql.DB
}

func NewQueries(db *sql.DB) *Queries {
	return &Queries{db: db}
}

type GetJobByIDRow struct {
	ID int64
}

func (q Queries) GetJobById(ctx context.Context, id int64) (GetJobByIDRow, error) {
	var row GetJobByIDRow
	return row, nil
}

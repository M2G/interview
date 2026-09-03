package infrastructure

import (
	"context"
	"interview/db"
)

type VideoJobDetails struct {
	ID int64
}

type SQLCRepo struct {
	queries *db.Queries
}

func NewSQLCRepo(q *db.Queries) *SQLCRepo {
	return &SQLCRepo{queries: q}
}

func (r *SQLCRepo) GetJobById(ctx context.Context, id int64) (VideoJobDetails, error) {
	row, err := r.queries.GetJobById(ctx, id)
	if err != nil {
		return VideoJobDetails{}, err
	}

	details := VideoJobDetails{
		ID: row.ID,
	}

	return details, nil
}

/*

func (r *SQLCRepository) // receiver
GetJobByID(ctx context.Context, id int64) // nom de la methode + param
(domain.VideoJobDetails, error) // valeur de retour
{ ... }

*/

/*
func main() {
	fmt.Println("Hello World from Go!")
}
*/

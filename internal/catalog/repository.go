package catalog

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CatalogRepository interface {
	List(ctx context.Context) ([]CatalogResponse, error)
}

type CatalogRepositoryImpl struct {
	Db *pgxpool.Pool
}

func NewCatalogRepository(db *pgxpool.Pool) CatalogRepository {
	return &CatalogRepositoryImpl{Db: db}
}

func (r *CatalogRepositoryImpl) List(ctx context.Context) ([]CatalogResponse, error) {
	rows, err := r.Db.Query(ctx, "SELECT * FROM categories")

	if err != nil {
		// return []CatalogResponse{}, err
		log.Fatalf("error on : %v", err)
	}

	defer rows.Close()

	responses := []CatalogResponse{}

	for rows.Next() {
		response := CatalogResponse{}

		rErr := rows.Scan(&response.Id, &response.Name)
		if rErr != nil {
			return responses, rErr
		}

		responses = append(responses, response)
	}

	return responses, nil
}

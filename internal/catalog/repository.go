package catalog

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CatalogRepository interface {
	List(ctx context.Context) ([]CatalogResponse, error)
	Get(ctx context.Context, id int) (CatalogResponse, error)
	Create(ctx context.Context, catalog CatalogResponse) (CatalogResponse, error)
	Delete(ctx context.Context, catalog CatalogResponse)
}

type CatalogRepositoryImpl struct {
	Db *pgxpool.Pool
}

func NewCatalogRepository(db *pgxpool.Pool) CatalogRepository {
	return &CatalogRepositoryImpl{Db: db}
}

var ErrNotFound error = errors.New("categories with that id not found")

func (r *CatalogRepositoryImpl) List(ctx context.Context) ([]CatalogResponse, error) {
	rows, err := r.Db.Query(ctx, "SELECT * FROM categories")

	if err != nil {
		return []CatalogResponse{}, fmt.Errorf("could not list : %w", err)
	}

	defer rows.Close()

	responses := []CatalogResponse{}

	for rows.Next() {
		response := CatalogResponse{}

		err := rows.Scan(&response.Id, &response.Name)
		if err != nil {
			return responses, fmt.Errorf("scan failed : %w", err)
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (r *CatalogRepositoryImpl) Get(ctx context.Context, id int) (CatalogResponse, error) {
	response := CatalogResponse{}
	rows := r.Db.QueryRow(ctx, "SELECT * FROM categories WHERE categoryId = $1", id).Scan(&response.Id, &response.Name)

	if errors.Is(rows, pgx.ErrNoRows) {
		return response, ErrNotFound
	}

	return response, nil
}

func (r *CatalogRepositoryImpl) Create(ctx context.Context, catalog CatalogResponse) (CatalogResponse, error) {
	tx, err := r.Db.Begin(ctx)

	if err != nil {
		return CatalogResponse{}, fmt.Errorf("cannot make tx : %w", err)
	}

	response := CatalogResponse{

		Name: catalog.Name,
	}

	rowErr := tx.QueryRow(ctx, "INSERT INTO categories (name) Values ($1) RETURNING categoryId", catalog.Name).Scan(&response.Id)

	if rowErr != nil {
		return CatalogResponse{}, rowErr
	}

	cErr := tx.Commit(ctx)

	if cErr != nil {
		return CatalogResponse{}, cErr
	}

	return response, nil
}

func (r *CatalogRepositoryImpl) Delete(ctx context.Context, catalog CatalogResponse) {
	tx, err := r.Db.Begin(ctx)

	if err != nil {
		return
	}

	_, dErr := tx.Exec(ctx, "DELETE FROM categories WHERE categoryId = $1", catalog.Id)

	if dErr != nil {
		return
	}

	if cErr := tx.Commit(ctx); cErr != nil {
		tx.Rollback(ctx)
	}
}

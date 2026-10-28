package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/RevanAdr/go-restful-api/internal/config"
	"github.com/RevanAdr/go-restful-api/internal/helper"
)

type CategoryRepository interface {
	GetById(id int) *helper.CategoryResponse
	GetAll() *[]helper.CategoryResponse
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) GetById(id int) *helper.CategoryResponse {
	conn := config.Load()
	defer conn.Close()

	var name string

	qErr := conn.QueryRow(context.Background(), "SELECT name FROM categories WHERE categoryId = $1", id).Scan(&name)

	if qErr != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", qErr)
		os.Exit(1)
	}

	response := helper.CategoryResponse{
		Id:   id,
		Name: name,
	}

	return &response
}

func (c *CategoryRepositoryImpl) GetAll() *[]helper.CategoryResponse {
	conn := config.Load()
	defer conn.Close()

	rows, rErrs := conn.Query(context.Background(), "SELECT * FROM categories")

	if rErrs != nil {
		fmt.Fprintf(os.Stderr, "Query failed : %v\n", rErrs)
	}

	var responses []helper.CategoryResponse

	for rows.Next() {
		var id int
		var name string

		if rows.Scan(&id, &name) != nil {
			fmt.Println("Scan failed")
		}

		response := helper.CategoryResponse{
			Id:   id,
			Name: name,
		}

		responses = append(responses, response)
	}

	return &responses
}
